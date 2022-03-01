import React from 'react';
import {Route, Routes} from "react-router-dom";
import Users from './pages/users';

const App = (): JSX.Element => <Routing />

const Routing: React.VFC<unknown> = () => (
  <>
      <Routes>
        <Route path="/" element={<Users />} />
      </Routes>
  </>
)


export default App;
