// Effective Error Handling
//
// Self Made Engineer — Advanced Go. Read theory below, then run:
//
//	go test ./advance-go/... -run Error -v
//	go test ./advance-go/... -race
//
// =============================================================================
// # 1. Go has no exceptions
// =============================================================================
// Functions signal failure by returning an error as the last value. Callers check:
//
//	if err != nil {
//	    return err
//	}
//
// Never ignore err with `_` unless you document why (rare).
//
// =============================================================================
// # 2. The error interface
// =============================================================================
//	type error interface {
//	    Error() string
//	}
//
// nil error means success. Any type with an Error() string method implements error.
//
// =============================================================================
// # 3. Creating errors
// =============================================================================
// - errors.New("message")           — simple static message (sentinel / package-level)
// - fmt.Errorf("format", args...)   — formatted message; use %w to wrap another error
//
// =============================================================================
// # 4. Sentinel errors
// =============================================================================
// Package-level variables compared with errors.Is:
//
//	var ErrNotFound = errors.New("not found")
//
// Good for stable, documented failure modes. Do not compare with == on wrapped errors.
//
// =============================================================================
// # 5. Custom error types
// =============================================================================
// Struct with Error() string — attach fields (code, op, path). Use errors.As to extract:
//
//	type OpError struct { Op string; Err error }
//	func (e *OpError) Error() string { return e.Op + ": " + e.Err.Error() }
//	func (e *OpError) Unwrap() error { return e.Err }  // optional, helps Is/As chain
//
// =============================================================================
// # 6. Wrapping (%w)
// =============================================================================
// Add context while preserving the chain:
//
//	return fmt.Errorf("read config %q: %w", path, err)
//
// errors.Is(wrapped, ErrNotFound) still works if the chain contains ErrNotFound.
// errors.As(wrapped, &target) finds a matching type in the chain.
//
// =============================================================================
// # 7. errors.Is vs errors.As
// =============================================================================
// - errors.Is(err, target)   — err == target OR any error in unwrap chain equals target
// - errors.As(err, &ptr)     — err or chain element is assignable to *ptr
//
// Prefer Is/As over err == ErrX or type assertions on wrapped errors.
//
// =============================================================================
// # 8. Multiple errors (Go 1.20+)
// =============================================================================
// errors.Join(err1, err2) — aggregate; useful when closing many resources.
//
// =============================================================================
// # 9. Patterns in production
// =============================================================================
// - Return early on err; avoid deep nesting.
// - Wrap at boundaries (HTTP handler, CLI, main) with operation context.
// - Log once at the top (main / middleware), not at every layer.
// - Map sentinel/custom errors to HTTP status in one place.
// - Do not panic for expected failures; panic for programmer bugs / init only.
//
// =============================================================================
// # Read
// =============================================================================
// - https://go.dev/blog/error-handling-and-go
// - https://pkg.go.dev/errors
// - Go Proverbs: https://go-proverbs.github.io/
// - 50 Shades — error gotchas: https://golang50shades.com/
//
// =============================================================================
// # Warm-up
// =============================================================================
// ParsePositive(s string) (int, error) — empty or non-positive → ErrInvalidInput.
//
// =============================================================================
// # Main exercise
// =============================================================================
// UserStore.Get(id): ErrNotFound if missing; wrap DB failures as *QueryError with Op field.
// Demonstrate errors.Is(ErrNotFound) and errors.As(*QueryError) in tests.
//
// =============================================================================
// # Large exercise — "Order checkout"
// =============================================================================
// Implement Checkout(cartID string) error that can return:
//   - ErrCartEmpty (sentinel)
//   - *PaymentError{Code, Message} for declined cards
//   - wrapped fmt.Errorf for inventory lookup failure
// Write table tests: Is for sentinel, As for PaymentError, message contains context for wrap.
//
package advancego

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// --- Sentinel errors (stable API contract) ---

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidInput  = errors.New("invalid input")
	ErrCartEmpty     = errors.New("cart is empty")
	ErrPaymentFailed = errors.New("payment failed")
)

// --- Custom error types ---

// QueryError describes a failed data access operation.
type QueryError struct {
	Op  string // e.g. "UserStore.Get"
	Err error
}

func (e *QueryError) Error() string {
	return fmt.Sprintf("%s: %v", e.Op, e.Err)
}

func (e *QueryError) Unwrap() error {
	return e.Err
}

// PaymentError is a domain error with a machine-readable code.
type PaymentError struct {
	Code    string
	Message string
}

func (e *PaymentError) Error() string {
	return fmt.Sprintf("payment %s: %s", e.Code, e.Message)
}

// ValidationError collects field-level problems (custom type + fields).
type ValidationError struct {
	Field   string
	Value   string
	Reason  string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("validation: %s %q: %s", e.Field, e.Value, e.Reason)
}

// --- Warm-up: ParsePositive ---

func ParsePositive(s string) (int, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("parse positive %q: %w", s, ErrInvalidInput)
	}
	n, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parse positive %q: %w", s, ErrInvalidInput)
	}
	if n <= 0 {
		return 0, fmt.Errorf("parse positive %q: %w", s, ErrInvalidInput)
	}
	return n, nil
}

// --- In-memory store demo (main exercise) ---

type User struct {
	ID   string
	Name string
}

// UserStore is a tiny in-memory repository.
type UserStore struct {
	users map[string]User
}

func NewUserStore(seed map[string]User) *UserStore {
	if seed == nil {
		seed = map[string]User{}
	}
	return &UserStore{users: seed}
}

func (s *UserStore) Get(id string) (User, error) {
	if strings.TrimSpace(id) == "" {
		return User{}, fmt.Errorf("user store get: %w", ErrInvalidInput)
	}
	u, ok := s.users[id]
	if !ok {
		return User{}, fmt.Errorf("user %q: %w", id, ErrNotFound)
	}
	return u, nil
}

// GetOrWrap simulates a DB failure wrapped as QueryError.
func (s *UserStore) GetOrWrap(id string, dbDown bool) (User, error) {
	if dbDown {
		return User{}, &QueryError{
			Op:  "UserStore.GetOrWrap",
			Err: errors.New("connection reset"),
		}
	}
	return s.Get(id)
}

// --- Large exercise: Checkout ---

type Cart struct {
	ID    string
	Items int
}

type Inventory struct {
	stock map[string]int
}

func NewInventory(stock map[string]int) *Inventory {
	return &Inventory{stock: stock}
}

func (inv *Inventory) Available(cartID string) (int, error) {
	n, ok := inv.stock[cartID]
	if !ok {
		return 0, fmt.Errorf("inventory lookup %q: %w", cartID, ErrNotFound)
	}
	return n, nil
}

type PaymentGateway struct {
	decline map[string]bool // cartID -> declined
}

func (pg *PaymentGateway) Charge(cartID string, amount int) error {
	if amount <= 0 {
		return ErrCartEmpty
	}
	if pg.decline != nil && pg.decline[cartID] {
		return &PaymentError{Code: "card_declined", Message: "insufficient funds"}
	}
	return nil
}

// Checkout runs inventory check then payment (composed errors).
func Checkout(cart Cart, inv *Inventory, pg *PaymentGateway) error {
	if cart.Items <= 0 {
		return ErrCartEmpty
	}
	stock, err := inv.Available(cart.ID)
	if err != nil {
		return fmt.Errorf("checkout %q: %w", cart.ID, err)
	}
	if stock < cart.Items {
		return fmt.Errorf("checkout %q: only %d in stock: %w", cart.ID, stock, ErrNotFound)
	}
	if err := pg.Charge(cart.ID, cart.Items); err != nil {
		return fmt.Errorf("checkout %q payment: %w", cart.ID, err)
	}
	return nil
}

// --- Helpers for handlers / CLI (map errors to user messages) ---

// HTTPStatus maps known errors to status codes (single place).
func HTTPStatus(err error) int {
	if err == nil {
		return 200
	}
	if errors.Is(err, ErrNotFound) {
		return 404
	}
	if errors.Is(err, ErrInvalidInput) {
		return 400
	}
	var pay *PaymentError
	if errors.As(err, &pay) {
		return 402
	}
	return 500
}

// UserMessage returns a safe message for clients (no internal details).
func UserMessage(err error) string {
	if err == nil {
		return ""
	}
	if errors.Is(err, ErrNotFound) {
		return "resource not found"
	}
	if errors.Is(err, ErrInvalidInput) {
		return "invalid request"
	}
	if errors.Is(err, ErrCartEmpty) {
		return "your cart is empty"
	}
	var pay *PaymentError
	if errors.As(err, &pay) {
		return "payment could not be completed"
	}
	return "something went wrong"
}

// JoinCloseErrors demonstrates errors.Join for cleanup.
func JoinCloseErrors(errs ...error) error {
	return errors.Join(errs...)
}
