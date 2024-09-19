import React from 'react';
import { Link } from 'react-router-dom';

const Navbar = ({ user, logout }) => {
  return (
    <nav className="bg-blue-600 text-white p-4">
      <div className="container mx-auto flex justify-between items-center">
        <Link to="/" className="text-2xl font-bold">E-Commerce</Link>
        <ul className="flex space-x-4">
          <li><Link to="/">Home</Link></li>
          <li><Link to="/products">Products</Link></li>
          {user ? (
            <>
              <li><Link to="/cart">Cart</Link></li>
              <li><Link to="/my-orders">My Orders</Link></li>
              <li><button onClick={logout}>Logout</button></li>
            </>
          ) : (
            <>
              <li><Link to="/login">Login</Link></li>
              <li><Link to="/register">Register</Link></li>
            </>
          )}
        </ul>
      </div>
    </nav>
  );
};

export default Navbar;