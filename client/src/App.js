import { Fragment } from 'react';
import { BrowserRouter as Router, Switch, Route } from 'react-router-dom';

import './App.css';
import 'bootstrap/dist/css/bootstrap.min.css';
import Layout from './hocs/Layout';
import Cart from './pages/Cart';
import Home from './pages/Home';
import Login from './pages/Login';
import Products from './pages/Products';
import Register from './pages/Register';

function App() {
  return (
    <div className="App">
      <Router>
        <Switch>
        <Fragment>
          <Layout>
            <Route exact path='/register' component={Register} />
            <Route exact path='/login' component={Login} />
            <Route exact path='/home' component={Home} />
            <Route exact path='/products' component={Products} />
            <Route exact path='/cart' component={Cart} />
          </Layout>
        </Fragment>
        </Switch>
      </Router>
    </div>
  );
}

export default App;
