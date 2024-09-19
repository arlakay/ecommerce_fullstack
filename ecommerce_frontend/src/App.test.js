import { render, screen } from '@testing-library/react';
import App from './App';

// Mocking the components to simplify the tests
jest.mock('./pages/Home/Home', () => () => <div>Home</div>);
jest.mock('./pages/Login/Login', () => () => <div>Login</div>);
jest.mock('./pages/Register/Register', () => () => <div>Register</div>);
jest.mock('./pages/ProductList/ProductList', () => () => <div>Product List</div>);
jest.mock('./pages/ProductDetail/ProductDetail', () => () => <div>Product Detail</div>);
jest.mock('./pages/Cart/Cart', () => () => <div>Cart</div>);
jest.mock('./pages/Order/Order', () => () => <div>Order</div>);
jest.mock('./pages/Order/UserOrderDetails', () => () => <div>User Orders</div>);
jest.mock('./components/Navbar/Navbar', () => () => <nav>Navbar</nav>);
jest.mock('./components/PrivateRoute/PrivateRoute', () => ({ children }) => <div>{children}</div>);


