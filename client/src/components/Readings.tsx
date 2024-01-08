import { useEffect, useState } from "react";
import { getReadings } from "../api/apiClient";
import { ReadingFull } from "../models/Reading";

const Index = () => {
  const [readings, setReadings] = useState<ReadingFull | null>(null);

  useEffect(() => {
    const fetchUser = async () => {
      return await getReadings();
    };

    fetchUser().then((users) => setReadings(users));
  }, []);

  console.log(readings);

  return <h1>Readings</h1>;
};

export default Index;
