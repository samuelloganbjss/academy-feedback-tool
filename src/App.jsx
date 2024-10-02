import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./Home.jsx";
import NoPage from "./NoPage.jsx";
import AdminDashboard from "./AdminDashboard.jsx";

const App = () => {


  return (
    <BrowserRouter>
      <Routes>
          <Route index element={<Home />} />
          <Route path="/admin" element={<AdminDashboard />} />
          <Route path="*" element={<NoPage />} />
      </Routes>
    </BrowserRouter>

  );
};

export default App;