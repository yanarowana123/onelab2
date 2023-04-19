package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/yanarowana123/onelab2/internal/models"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type TransactionRepository struct {
	baseUrl   string
	client    http.Client
	secretKey string
}

func NewTransactionRepository(baseUrl string, timeout time.Duration, secretKey string) *TransactionRepository {
	return &TransactionRepository{
		baseUrl:   baseUrl,
		client:    http.Client{Timeout: timeout},
		secretKey: secretKey,
	}
}

func (r *TransactionRepository) Create(ctx context.Context, createTransactionRequest models.CreateTransactionRequest) error {
	jsonBody, err := json.Marshal(createTransactionRequest)
	if err != nil {
		return err
	}

	bodyReader := bytes.NewReader(jsonBody)

	requestURL := fmt.Sprintf("%s/transaction", r.baseUrl)
	request, err := r.newRequest(ctx, http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return err
	}

	response, err := r.client.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		return errors.New("transaction error")
	}

	return nil
}
func (r *TransactionRepository) GetSumByBookID(ctx context.Context, bookID uuid.UUID) (float64, error) {
	requestURL := fmt.Sprintf("%s/sum/book/%s", r.baseUrl, bookID)
	request, err := r.newRequest(ctx, http.MethodGet, requestURL, nil)
	if err != nil {
		return 0, err
	}

	response, err := r.client.Do(request)
	if err != nil {
		return 0, err
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}
	sum, err := strconv.ParseFloat(strings.TrimSpace(string(body)), 64)
	return sum, nil
}

func (r *TransactionRepository) newRequest(ctx context.Context, method string, url string, bodyReader io.Reader) (*http.Request, error) {
	request, err := http.NewRequestWithContext(ctx, method, url, bodyReader)
	request.Header.Set("Secret-Key", r.secretKey)
	return request, err
}
