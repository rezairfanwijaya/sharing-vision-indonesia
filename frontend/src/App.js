import AddNew from './pages/AddNew';
import AllPosts from './pages/AllPost';
import Index from './pages/Index';
import Preview from './pages/Preview';
import './style/style.css'
import { Routes, Route } from 'react-router-dom';

function App() {
  return (
    <div className="App px-10 py-5">
      <Routes>
        <Route
          path={"/"}
          exact
          element={<Index />}
        >
          <Route
            path={"/posts"}
            exact
            element={<AllPosts />}
          />
          <Route
            path={"/new"}
            exact
            element={<AddNew />}
          />
          <Route
            path={"/preview"}
            exact
            element={<Preview />}
          />
        </Route>
      </Routes>
    </div>
  );
}

export default App;
