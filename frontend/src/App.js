import AddNew from './pages/AddNew';
import AllPosts from './pages/AllPost';
import Index from './pages/Index';
import Preview from './pages/Preview';
import './style/style.css'
import { Routes, Route } from 'react-router-dom';
//theme
import "primereact/resources/themes/lara-light-indigo/theme.css";

//core
import "primereact/resources/primereact.min.css";
import "primeicons/primeicons.css";
import EditArticle from './pages/EditArticle';

function App() {
  return (
    <div className="App px-10 py-5">
      <Routes>
        <Route
          path={""}
          exact
          element={<Index />}
        >
          <Route
            path={"/"}
            exact
            element={<AllPosts />}
          />
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
          <Route
            path={"/article/:id"}
            exact
            element={<EditArticle/>}
          />
        </Route>
      </Routes>
    </div>
  );
}

export default App;
