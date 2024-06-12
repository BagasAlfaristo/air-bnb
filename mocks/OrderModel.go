// Code generated by mockery v2.43.2. DO NOT EDIT.

package mocks

import (
	order "PetPalApp/features/order"

	mock "github.com/stretchr/testify/mock"
)

// OrderModel is an autogenerated mock type for the OrderModel type
type OrderModel struct {
	mock.Mock
}

// CreateOrder provides a mock function with given fields: _a0
func (_m *OrderModel) CreateOrder(_a0 order.Order) (order.Order, error) {
	ret := _m.Called(_a0)

	if len(ret) == 0 {
		panic("no return value specified for CreateOrder")
	}

	var r0 order.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(order.Order) (order.Order, error)); ok {
		return rf(_a0)
	}
	if rf, ok := ret.Get(0).(func(order.Order) order.Order); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(order.Order)
	}

	if rf, ok := ret.Get(1).(func(order.Order) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByID provides a mock function with given fields: orderID
func (_m *OrderModel) GetOrderByID(orderID uint) (*order.Order, error) {
	ret := _m.Called(orderID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderByID")
	}

	var r0 *order.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*order.Order, error)); ok {
		return rf(orderID)
	}
	if rf, ok := ret.Get(0).(func(uint) *order.Order); ok {
		r0 = rf(orderID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(orderID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderByInvoiceID provides a mock function with given fields: invoiceID
func (_m *OrderModel) GetOrderByInvoiceID(invoiceID string) (*order.Order, error) {
	ret := _m.Called(invoiceID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrderByInvoiceID")
	}

	var r0 *order.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (*order.Order, error)); ok {
		return rf(invoiceID)
	}
	if rf, ok := ret.Get(0).(func(string) *order.Order); ok {
		r0 = rf(invoiceID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(invoiceID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrdersByUserID provides a mock function with given fields: userID
func (_m *OrderModel) GetOrdersByUserID(userID uint) ([]order.Order, error) {
	ret := _m.Called(userID)

	if len(ret) == 0 {
		panic("no return value specified for GetOrdersByUserID")
	}

	var r0 []order.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) ([]order.Order, error)); ok {
		return rf(userID)
	}
	if rf, ok := ret.Get(0).(func(uint) []order.Order); ok {
		r0 = rf(userID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]order.Order)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(userID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetProductByID provides a mock function with given fields: productID
func (_m *OrderModel) GetProductByID(productID uint) (*order.Product, error) {
	ret := _m.Called(productID)

	if len(ret) == 0 {
		panic("no return value specified for GetProductByID")
	}

	var r0 *order.Product
	var r1 error
	if rf, ok := ret.Get(0).(func(uint) (*order.Product, error)); ok {
		return rf(productID)
	}
	if rf, ok := ret.Get(0).(func(uint) *order.Product); ok {
		r0 = rf(productID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*order.Product)
		}
	}

	if rf, ok := ret.Get(1).(func(uint) error); ok {
		r1 = rf(productID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// UpdateOrder provides a mock function with given fields: orderID, _a1
func (_m *OrderModel) UpdateOrder(orderID uint, _a1 order.Order) (order.Order, error) {
	ret := _m.Called(orderID, _a1)

	if len(ret) == 0 {
		panic("no return value specified for UpdateOrder")
	}

	var r0 order.Order
	var r1 error
	if rf, ok := ret.Get(0).(func(uint, order.Order) (order.Order, error)); ok {
		return rf(orderID, _a1)
	}
	if rf, ok := ret.Get(0).(func(uint, order.Order) order.Order); ok {
		r0 = rf(orderID, _a1)
	} else {
		r0 = ret.Get(0).(order.Order)
	}

	if rf, ok := ret.Get(1).(func(uint, order.Order) error); ok {
		r1 = rf(orderID, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NewOrderModel creates a new instance of OrderModel. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewOrderModel(t interface {
	mock.TestingT
	Cleanup(func())
}) *OrderModel {
	mock := &OrderModel{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
