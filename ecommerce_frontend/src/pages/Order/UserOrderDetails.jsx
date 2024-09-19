import React, { useState, useEffect, useCallback } from 'react';
import { useNavigate } from 'react-router-dom';

const UserOrderDetails = ({ user, token, logout }) => {
    const [orders, setOrders] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const [selectedOrder, setSelectedOrder] = useState(null);
    const navigate = useNavigate();

    const handleSessionExpired = useCallback(() => {
        logout();
        navigate('/login', { state: { message: 'Your session has expired. Please log in again.' } });
    }, [logout, navigate]);

    const fetchOrders = useCallback(async () => {
        try {
            const response = await fetch(`http://localhost:8080/orders/user/${user.id}`, {
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.status === 401) {
                handleSessionExpired();
                return;
            }

            if (!response.ok) {
                throw new Error('Failed to fetch orders');
            }

            const data = await response.json();
            setOrders(data);
            setLoading(false);
        } catch (err) {
            setError(err.message);
            setLoading(false);
        }
    }, [user.id, token, handleSessionExpired]);

    useEffect(() => {
        if (!token) {
            navigate('/login');
            return;
        }
        fetchOrders();
    }, [token, fetchOrders, navigate]);

    const selectOrder = (order) => {
        setSelectedOrder(order);
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>Error: {error}</div>;

    return (
        <div className="container mx-auto p-4">
            <h1 className="text-2xl font-bold mb-4">Your Orders</h1>

            <div className="grid grid-cols-1 md:grid-cols-2 gap-4">
                <div className="order-list">
                    {orders.map(order => (
                        <div
                            key={order.id}
                            className="bg-white p-4 rounded shadow mb-4 cursor-pointer hover:bg-gray-100"
                            onClick={() => selectOrder(order)}
                        >
                            <p>Order ID: {order.id}</p>
                            <p>Date: {new Date(order.created_at).toLocaleDateString()}</p>
                            <p>Total: Rp {order.total.toFixed(2)}</p>
                            <p>Status: {order.status}</p>
                        </div>
                    ))}
                </div>

                <div className="order-details">
                    {selectedOrder && (
                        <div className="bg-white p-4 rounded shadow">
                            <h2 className="text-xl font-bold mb-2">Order Details</h2>
                            <p>Order ID: {selectedOrder.id}</p>
                            <p>Date: {new Date(selectedOrder.created_at).toLocaleDateString()}</p>
                            <p>Status: {selectedOrder.status}</p>
                            <h3 className="font-bold mt-4 mb-2">Items:</h3>
                            <ul>
                                {selectedOrder.order_items.map(item => (
                                    <li key={item.id} className="mb-2">
                                        <p>Product ID: {item.product_id}</p>
                                        <p>Quantity: {item.quantity}</p>
                                        <p>Price: Rp {item.price.toFixed(2)}</p>
                                    </li>
                                ))}
                            </ul>
                            <p className="font-bold mt-4">Total: Rp {selectedOrder.total.toFixed(2)}</p>
                        </div>
                    )}
                </div>
            </div>
        </div>
    );
};

export default UserOrderDetails;