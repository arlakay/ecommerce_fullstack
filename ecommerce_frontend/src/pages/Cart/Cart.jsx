import React, { useState, useEffect } from 'react';
import { useNavigate } from 'react-router-dom';

const Cart = ({ user, token, logout }) => {
    const [cartItems, setCartItems] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const navigate = useNavigate();

    useEffect(() => {
        if (!token) {
            navigate('/login');
            return;
        }

        const fetchCart = async () => {
            try {
                const response = await fetch(`http://localhost:8080/cart/user/${user.id}`, {
                    headers: {
                        'Authorization': `Bearer ${token}`
                    }
                });

                if (response.status === 401) {
                    logout();
                    navigate('/login', { state: { message: 'Your session has expired. Please log in again.' } });
                    return;
                }

                if (!response.ok) {
                    throw new Error('Failed to fetch cart');
                }
                const data = await response.json();
                setCartItems(data);
                setLoading(false);
            } catch (err) {
                setError(err.message);
                setLoading(false);
            }
        };

        fetchCart();
    }, [user, token, navigate, logout]);

    // Helper function to update item quantity
    const updateItemQuantity = async (itemId, productId, quantity) => {
        if (quantity < 1) return; // Prevent quantity from going below 1

        try {
            const response = await fetch(`http://localhost:8080/cart/${itemId}`, {
                method: 'PUT',
                headers: {
                    'Authorization': `Bearer ${token}`,
                    'Content-Type': 'application/json'
                },
                body: JSON.stringify({
                    user_id: user.id,
                    product_id: productId, // Assuming itemId corresponds to product_id
                    quantity: quantity
                })
            });

            if (response.status === 401) {
                logout();
                navigate('/login', { state: { message: 'Your session has expired. Please log in again.' } });
                return;
            }

            if (!response.ok) {
                throw new Error('Failed to update quantity');
            }

            // Update cart items with the new quantity
            setCartItems(prevItems =>
                prevItems.map(item =>
                    item.id === itemId ? { ...item, quantity } : item
                )
            );

        } catch (err) {
            setError(err.message);
        }
    };

    const handleIncreaseQuantity = (itemId, productId, currentQuantity, stock) => {
        if (currentQuantity < stock) {
            updateItemQuantity(itemId, productId, currentQuantity + 1);
        }
    };

    const handleDecreaseQuantity = (itemId, productId, currentQuantity) => {
        if (currentQuantity > 1) {
            updateItemQuantity(itemId, productId, currentQuantity - 1);
        }
    };

    // Helper function to remove an item from the cart
    const handleRemoveItem = async (itemId) => {
        try {
            const response = await fetch(`http://localhost:8080/cart/${itemId}`, {
                method: 'DELETE',
                headers: {
                    'Authorization': `Bearer ${token}`
                }
            });

            if (response.status === 401) {
                logout();
                navigate('/login', { state: { message: 'Your session has expired. Please log in again.' } });
                return;
            }

            if (!response.ok) {
                throw new Error('Failed to remove item');
            }
            // Update cart items after removal
            setCartItems(cartItems.filter(item => item.id !== itemId));
        } catch (err) {
            setError(err.message);
        }
    };

    // Helper function to calculate the total price of the cart
    const calculateTotal = () => {
        return cartItems.reduce((total, item) => total + item.product.price * item.quantity, 0);
    };

    const handleProceedToCheckout = () => {
        navigate('/order', { state: { cartItems } });
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;

    return (
        <div>
            <h2 className="text-2xl font-bold mb-4">Your Cart</h2>
            {cartItems.length === 0 ? (
                <p>Your cart is empty</p>
            ) : (
                <div>
                    <ul className="list-disc pl-5">
                        {cartItems.map(item => {
                            const price = item && typeof item.product.price === 'number' ? item.product.price.toFixed(2) : 'N/A';

                            return (
                                <li key={item.id} className="mb-4 flex items-center">
                                    <img
                                        src={item.image} // Assuming each item has an image URL
                                        alt={item.name}
                                        className="w-16 h-16 object-cover mr-4" />
                                    <div className="flex-grow">
                                        <h3 className="text-lg font-semibold">{item.name}</h3>
                                        <p>Price: Rp {price}</p>
                                        {/* <p>Quantity: {item.quantity}</p> */}
                                        <div className="flex items-center">
                                            <button
                                                onClick={() => handleDecreaseQuantity(item.id, item.product_id, item.quantity)}
                                                className="px-2 py-1 bg-gray-300 rounded"
                                            >
                                                -
                                            </button>
                                            <span className="mx-2">{item.quantity}</span>
                                            <button
                                                onClick={() => handleIncreaseQuantity(item.id, item.product_id, item.quantity, item.product.stock)}
                                                className="px-2 py-1 bg-gray-300 rounded"
                                            >
                                                +
                                            </button>
                                        </div>
                                    </div>
                                    <button
                                        onClick={() => handleRemoveItem(item.id)}
                                        className="ml-4 px-4 py-2 bg-red-500 text-white rounded hover:bg-red-600"
                                    >
                                        Remove
                                    </button>
                                </li>
                            );
                        })}
                    </ul>
                    <div className="mt-4 flex justify-between items-center">
                        <h3 className="text-xl font-bold">Total: Rp{calculateTotal().toFixed(2)}</h3>
                        <button
                            onClick={handleProceedToCheckout}
                            className="px-4 py-2 bg-green-500 text-white rounded hover:bg-green-600"
                        >
                            Proceed to Checkout
                        </button>
                    </div>
                </div>
            )}
        </div>
    );
};

export default Cart;
