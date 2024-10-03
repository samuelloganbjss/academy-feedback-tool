import { BrowserRouter, Routes, Route } from "react-router-dom";
import Home from "./Home.jsx";
import NoPage from "./NoPage.jsx";
import AdminDashboard from "./AdminDashboard.jsx";
import Banner from "./Banner.jsx";
import 'bootstrap/dist/css/bootstrap.min.css';
import "@fontsource/poppins";
import "./App.css"

const App = () => {

  return (
    <BrowserRouter>
    <Banner />
      <Routes>
        
          <Route index element={<Home />} />
          <Route path="/admin" element={<AdminDashboard />} />
          <Route path="*" element={<NoPage />} />
      </Routes>
    </BrowserRouter>

  );
};

export default App;