import React, { useState, useEffect } from 'react';
import { useParams, useNavigate } from 'react-router-dom';

const ProductDetail = ({ user, token }) => {

    const [product, setProduct] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');
    const { id } = useParams();
    const navigate = useNavigate();

    useEffect(() => {
        const fetchProduct = async () => {
            try {
                const response = await fetch(`http://localhost:8080/products/${id}`);
                if (!response.ok) {
                    throw new Error('Failed to fetch product');
                }
                const data = await response.json();
                setProduct(data);
                setLoading(false);
            } catch (err) {
                setError(err.message);
                setLoading(false);
            }
        };

        fetchProduct();
    }, [id]);

    const handleAddToCart = async () => {
        if (!user) {
            navigate('/login');
            return;
        }

        const requestBody = {
            user_id: user.id,
            product_id: product.id,
            quantity: 1,
        };

        try {
            const response = await fetch('http://localhost:8080/cart', {
                method: 'POST',
                headers: {
                    'Content-Type': 'application/json',
                    'Authorization': `Bearer ${token}`
                },
                body: JSON.stringify(requestBody),
            });

            if (!response.ok) {
                throw new Error('Failed to add product to cart');
            }

            navigate('/cart');
        } catch (err) {
            setError(err.message);
        }
    };

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;
    if (!product) return <div>Product not found</div>;

    return (
        <div>
            <h2 className="text-2xl font-bold mb-4">{product.name}</h2>
            <img src={product.image_url || `https://placehold.co/400`} alt={product.name} className="w-full max-w-md mb-4" />
            <p className="mb-4">{product.description}</p>
            <p className="text-xl mb-4">${product.price}</p>
            <button
                onClick={handleAddToCart}
                className="bg-blue-500 text-white py-2 px-4 rounded hover:bg-blue-600"
            >
                Add to Cart
            </button>
        </div>
    );
};

export default ProductDetail;