package advancego

import (
	"errors"
	"testing"
)

func TestParsePositive(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		want    int
		wantErr error
	}{
		{"ok", "42", 42, nil},
		{"empty", "", 0, ErrInvalidInput},
		{"zero", "0", 0, ErrInvalidInput},
		{"negative", "-3", 0, ErrInvalidInput},
		{"bad", "abc", 0, ErrInvalidInput},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParsePositive(tt.in)
			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Fatalf("errors.Is: got %v want %v", err, tt.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			if got != tt.want {
				t.Fatalf("got %d want %d", got, tt.want)
			}
		})
	}
}

func TestUserStore_Get(t *testing.T) {
	store := NewUserStore(map[string]User{"u1": {ID: "u1", Name: "Ada"}})

	_, err := store.Get("")
	if !errors.Is(err, ErrInvalidInput) {
		t.Fatalf("empty id: %v", err)
	}

	_, err = store.Get("missing")
	if !errors.Is(err, ErrNotFound) {
		t.Fatalf("missing: %v", err)
	}

	u, err := store.Get("u1")
	if err != nil || u.Name != "Ada" {
		t.Fatalf("got %+v err %v", u, err)
	}
}

func TestUserStore_GetOrWrap_AsQueryError(t *testing.T) {
	store := NewUserStore(nil)
	_, err := store.GetOrWrap("u1", true)
	var qe *QueryError
	if !errors.As(err, &qe) {
		t.Fatalf("expected *QueryError, got %T %v", err, err)
	}
	if qe.Op != "UserStore.GetOrWrap" {
		t.Fatalf("Op = %q", qe.Op)
	}
}

func TestCheckout(t *testing.T) {
	inv := NewInventory(map[string]int{"c1": 5})

	tests := []struct {
		name       string
		cart       Cart
		decline    bool
		wantIs     error
		wantAsPay  bool
	}{
		{
			name:   "empty cart",
			cart:   Cart{ID: "c1", Items: 0},
			wantIs: ErrCartEmpty,
		},
		{
			name:   "success",
			cart:   Cart{ID: "c1", Items: 2},
		},
		{
			name:      "payment declined",
			cart:      Cart{ID: "c1", Items: 1},
			decline:   true,
			wantAsPay: true,
		},
		{
			name:   "unknown cart inventory",
			cart:   Cart{ID: "unknown", Items: 1},
			wantIs: ErrNotFound,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gw := &PaymentGateway{decline: map[string]bool{}}
			if tt.decline {
				gw.decline[tt.cart.ID] = true
			}
			err := Checkout(tt.cart, inv, gw)
			if tt.wantIs != nil {
				if !errors.Is(err, tt.wantIs) {
					t.Fatalf("Is: %v", err)
				}
				return
			}
			if tt.wantAsPay {
				var pe *PaymentError
				if !errors.As(err, &pe) {
					t.Fatalf("As PaymentError: %v", err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestHTTPStatus_and_UserMessage(t *testing.T) {
	if HTTPStatus(ErrNotFound) != 404 {
		t.Fatal("404")
	}
	if HTTPStatus(&PaymentError{Code: "x", Message: "y"}) != 402 {
		t.Fatal("402")
	}
	if UserMessage(ErrCartEmpty) != "your cart is empty" {
		t.Fatal(UserMessage(ErrCartEmpty))
	}
}

func TestJoinCloseErrors(t *testing.T) {
	e1 := errors.New("close a")
	e2 := errors.New("close b")
	j := JoinCloseErrors(e1, e2)
	if j == nil {
		t.Fatal("expected joined error")
	}
	if !errors.Is(j, e1) || !errors.Is(j, e2) {
		t.Fatalf("%v", j)
	}
}
