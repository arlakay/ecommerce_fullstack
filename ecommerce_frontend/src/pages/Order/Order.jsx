import React, { useState, useEffect } from 'react';
import { useNavigate, useLocation } from 'react-router-dom';

const Order = ({ user, token, logout }) => {
    const [orderStatus, setOrderStatus] = useState('');
    const [error, setError] = useState('');
    const navigate = useNavigate();
    const location = useLocation();
    const cartItems = location.state?.cartItems || [];

    useEffect(() => {
        if (!token) {
            navigate('/login');
        }
    }, [token, navigate]);

    const createOrder = async () => {
        try {
            const orderItems = cartItems.map(item => ({
                product_id: item.product.id,
                quantity: item.quantity,
                price: item.product.price
            }));

            const total = cartItems.reduce((sum, item) => sum + item.product.price * item.quantity, 0);

            const response = await fetch('http://localhost:8080/orders', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify({
                    user_id: user.id,
                    order_items: orderItems,
                    total: total,
                    status: 'pending'
                })
            });

            if (response.status === 401) {
                logout();
                navigate('/login', { state: { message: 'Your session has expired. Please log in again.' } });
                return;
            }

            if (!response.ok) {
                throw new Error('Failed to create order');
            }

            const data = await response.json();
            setOrderStatus('Order placed successfully!');
            // You might want to clear the cart or redirect to an order confirmation page here
        } catch (err) {
            setError(err.message);
        }
    };

    return (
        <div className="max-w-2xl mx-auto mt-8 p-6 bg-white rounded-lg shadow-md">
            <h2 className="text-2xl font-bold mb-6">Confirm Your Order</h2>

            {cartItems.length === 0 ? (
                <p>Your cart is empty. Please add items before checking out.</p>
            ) : (
                <>
                    <ul className="mb-6">
                        {cartItems.map(item => (
                            <li key={item.id} className="flex justify-between items-center mb-2">
                                <span>{item.product.name} (x{item.quantity})</span>
                                <span>Rp {(item.product.price * item.quantity).toFixed(2)}</span>
                            </li>
                        ))}
                    </ul>

                    <div className="flex justify-between items-center font-bold text-lg mb-6">
                        <span>Total:</span>
                        <span>Rp {cartItems.reduce((sum, item) => sum + item.product.price * item.quantity, 0).toFixed(2)}</span>
                    </div>

                    <button
                        onClick={createOrder}
                        className="w-full py-2 px-4 bg-blue-500 text-white rounded hover:bg-blue-600 transition duration-200"
                    >
                        Place Order
                    </button>

                    {orderStatus && <p className="mt-4 text-green-600">{orderStatus}</p>}
                    {error && <p className="mt-4 text-red-600">{error}</p>}
                </>
            )}
        </div>
    );
};

export default Order;