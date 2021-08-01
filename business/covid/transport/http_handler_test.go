package transport_test

import (
	"covid-19/business/covid/model"
	"covid-19/business/covid/transport"
	mock_covid "covid-19/test"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"gotest.tools/assert"
)

func TestCovidHTTPHandler_GetCovidSummary_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mSrc := mock_covid.NewMockService(ctrl)

	ts := httptest.NewServer(transport.NewCovidHTTPHandler(mSrc))
	defer ts.Close()

	// case call api success
	mSrc.EXPECT().CovidSummary().Return(model.Summary{}, nil)
	resp, err := http.Get(fmt.Sprintf("%s/covid/summary", ts.URL))

	assert.Assert(t, resp.Body != nil)
	assert.Assert(t, err == nil)

}

func TestCovidHTTPHandler_GetCovidSummary_Fail(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mSrc := mock_covid.NewMockService(ctrl)

	ts := httptest.NewServer(transport.NewCovidHTTPHandler(mSrc))
	defer ts.Close()

	// case call api fail
	mSrc.EXPECT().CovidSummary().Return(model.Summary{}, errors.New("something wrong"))
	resp, err := http.Get(fmt.Sprintf("%s/covid/summary", ts.URL))
	assert.Assert(t, err == nil)
	assert.Assert(t, resp.StatusCode == http.StatusInternalServerError)

}
