import { Route, Routes } from "react-router-dom";
import "./App.css";
import Index from "./components/Index";
import Readings from "./components/Readings";

function App() {
  return (
    <>
      <Routes>
        <Route path="/" element={<Index />} />
        <Route path="/readings" element={<Readings />} />
      </Routes>
    </>
  );
}

export default App;
