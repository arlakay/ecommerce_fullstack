import logo from './logo.svg';
import './App.css';
import { Link, Route, Router, Routes } from 'react-router-dom';
import { useEffect, useState } from 'react';
import Home from './pages/Home/Home';
import Login from './pages/Login/Login';
import Register from './pages/Register/Register';
import ProductList from './pages/ProductList/ProductList';
import ProductDetail from './pages/ProductDetail/ProductDetail';
import PrivateRoute from './components/PrivateRoute/PrivateRoute';
import Cart from './pages/Cart/Cart';
import Navbar from './components/Navbar/Navbar';
import OrderPage from './pages/Order/Order';
import Order from './pages/Order/Order';
import UserOrderDetails from './pages/Order/UserOrderDetails';

function App() {
  const [user, setUser] = useState(null);
  const [token, setToken] = useState(null);

  useEffect(() => {
    const storedUser = localStorage.getItem('user');
    if (storedUser) {
      setUser(JSON.parse(storedUser));
    }
  }, []);

  const login = (token, userData) => {
    setToken(token)
    setUser(userData);
    localStorage.setItem('token', JSON.stringify(token));
    localStorage.setItem('user', JSON.stringify(userData));
  };

  const logout = () => {
    setToken(null);
    setUser(null);
    localStorage.removeItem('token');
    localStorage.removeItem('user');
  };

  return (
    <div className="min-h-screen flex flex-col">
      <header className="bg-blue-600 text-white p-4">
        <Navbar user={user} logout={logout} />
      </header>

      <main className="flex-grow container mx-auto mt-4 p-4">
        <Routes>
          <Route path="/" element={<Home />} />
          <Route path="/login" element={<Login login={login} />} />
          <Route path="/register" element={<Register login={login} />} />
          <Route path="/products" element={<ProductList />} />
          <Route path="/products/:id" element={<ProductDetail user={user} token={token} />} />
          <Route
            path="/cart"
            element={
              <PrivateRoute user={user}>
                <Cart user={user} token={token} logout={logout} />
              </PrivateRoute>
            }
          />
          <Route
            path="/order"
            element={
              <PrivateRoute user={user}>
                <Order
                  user={user} token={token} logout={logout}
                />
              </PrivateRoute>
            }
          />
          <Route
            path="/my-orders"
            element={
              <PrivateRoute user={user}>
                <UserOrderDetails user={user} token={token} logout={logout} />
              </PrivateRoute>
            }
          />
        </Routes>
      </main>

      <footer className="bg-gray-200 p-4 mt-8">
        <div className="container mx-auto text-center">
          © 2024 E-Commerce. All rights reserved.
        </div>
      </footer>
    </div>
  );
}

export default App;
