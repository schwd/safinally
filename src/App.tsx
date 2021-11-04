import React, { Fragment, useEffect, useState } from 'react';
import { BrowserRouter as Router, Switch, Route } from "react-router-dom";
import Navbar from "./components/Navbar";
import Refers from "./components/Refers";
import ReferCreate from "./components/ReferCreate";
import SignIn from "./components/SignIn";
import Home from './components/Home';

function App() {
  const [token, setToken] = useState<string>("");

  useEffect(() => {
    const getToken = localStorage.getItem("token");
    if (getToken) {
      setToken(getToken);
    }
  }, []);

  if (!token) {
    return <SignIn />
  }

  return (
    <div>
      <Router>
        {token && (
          <Fragment>
            <Navbar/>
            <Switch>
              <Route exact path="/" component={Home} />
              <Route exact path="/Refers" component={Refers} />
              <Route exact path="/link/ReferCreate" component={ReferCreate} />
            </Switch>
          </Fragment>
        )}
      </Router>
    </div>

  );
}

export default App;