import React, { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { CurrencyFormatter } from '../../utils/CurrencyFormatter';

const Home = () => {

    const [featuredProducts, setFeaturedProducts] = useState([]);
    const [categories, setCategories] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState('');

    useEffect(() => {
        const fetchFeaturedProducts = async () => {
            try {
                const response = await fetch('http://localhost:8080/products');
                if (!response.ok) {
                    throw new Error('Failed to fetch featured products');
                }
                const data = await response.json();
                setFeaturedProducts(data);
                setLoading(false);
            } catch (err) {
                setError(err.message);
                setLoading(false);
            }
        };

        fetchFeaturedProducts();
    }, []);

    useEffect(() => {
        const fetchCategories = async () => {
            try {
                const response = await fetch('http://localhost:8080/categories');
                if (!response.ok) {
                    throw new Error('Failed to fetch featured products');
                }
                const data = await response.json();
                setCategories(data);
                setLoading(false);
            } catch (err) {
                setError(err.message);
                setLoading(false);
            }
        };

        fetchCategories();
    }, []);

    if (loading) return <div>Loading...</div>;
    if (error) return <div>{error}</div>;

    return (
        <>
            <div className="home-page">
                <section className="categories my-12">
                    <h2 className="text-2xl font-bold mb-6">Shop by Category</h2>
                    <div className="grid grid-cols-2 sm:grid-cols-3 md:grid-cols-4 gap-4">
                        {categories.map((category) => (
                            <div key={category} className="bg-gray-200 p-4 rounded-lg text-center">
                                <Link to={`/products?category=${category}`} className="text-lg font-semibold hover:underline">
                                    {category.name}
                                </Link>
                            </div>
                        ))}
                    </div>
                </section>

                <section className="featured-products my-12">
                    <h2 className="text-2xl font-bold mb-6">Featured Products</h2>
                    <div className="grid grid-cols-1 sm:grid-cols-2 md:grid-cols-3 lg:grid-cols-4 gap-6">
                        {featuredProducts.map((product) => (
                            <div key={product.id} className="bg-white shadow rounded-lg overflow-hidden">
                                <img src={product.image_url || `https://placehold.co/400`} alt={product.name} className="w-full h-48 object-cover" />
                                <div className="p-4">
                                    <h3 className="text-lg font-semibold mb-2">{product.name}</h3>
                                    <p className="text-gray-600 mb-2"><CurrencyFormatter amount={product.price} /></p>
                                    <Link to={`/products/${product.id}`} className="text-blue-500 hover:underline">
                                        View Details
                                    </Link>
                                </div>
                            </div>
                        ))}
                    </div>
                </section>
            </div>
        </>
    );
};

export default Home;
