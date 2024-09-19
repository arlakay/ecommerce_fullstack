const API_URL = 'http://localhost:8080'; // Replace with your backend URL

const handleResponse = async (response) => {
    if (!response.ok) {
        const error = await response.text();
        throw new Error(error);
    }
    return response.json();
};

const getAuthHeader = () => {
    const user = JSON.parse(localStorage.getItem('user'));
    return user && user.token ? { 'Authorization': `Bearer ${user.token}` } : {};
};

const api = {
    login: (email, password) =>
        fetch(`${API_URL}/login`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ email, password }),
        }).then(handleResponse),

    register: (name, email, password) =>
        fetch(`${API_URL}/register`, {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ name, email, password }),
        }).then(handleResponse),

    getProducts: () =>
        fetch(`${API_URL}/products`, {
            headers: getAuthHeader(),
        }).then(handleResponse),

    getProduct: (id) =>
        fetch(`${API_URL}/products/${id}`, {
            headers: getAuthHeader(),
        }).then(handleResponse),

    getCart: () =>
        fetch(`${API_URL}/cart`, {
            headers: getAuthHeader(),
        }).then(handleResponse),

    addToCart: (productId, quantity) =>
        fetch(`${API_URL}/cart`, {
            method: 'POST',
            headers: {
                ...getAuthHeader(),
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ productId, quantity }),
        }).then(handleResponse),

    updateCartItem: (id, quantity) =>
        fetch(`${API_URL}/cart/${id}`, {
            method: 'PUT',
            headers: {
                ...getAuthHeader(),
                'Content-Type': 'application/json',
            },
            body: JSON.stringify({ quantity }),
        }).then(handleResponse),

    removeFromCart: (id) =>
        fetch(`${API_URL}/cart/${id}`, {
            method: 'DELETE',
            headers: getAuthHeader(),
        }).then(handleResponse),

    getOrders: () =>
        fetch(`${API_URL}/orders`, {
            headers: getAuthHeader(),
        }).then(handleResponse),

    createOrder: (orderData) =>
        fetch(`${API_URL}/orders`, {
            method: 'POST',
            headers: {
                ...getAuthHeader(),
                'Content-Type': 'application/json',
            },
            body: JSON.stringify(orderData),
        }).then(handleResponse),
};

export default api;