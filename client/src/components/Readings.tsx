import { useEffect, useState } from "react";
import { getReadings } from "../api/apiClient";
import { ReadingFull } from "../models/Reading";

const Index = () => {
  const [readings, setReadings] = useState<ReadingFull[]>([]);

  useEffect(() => {
    const fetchReading = async () => {
      return await getReadings();
    };

    fetchReading().then((readings) => setReadings(readings));
  }, []);

  return (
    <div>
      <h1>Readings</h1>
      {readings ? (
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Temperature</th>
              <th>Humidity</th>
              <th>Timestamp</th>
            </tr>
          </thead>
          <tbody>
            {readings.map((reading) => (
              <tr key={reading.ID}>
                <td>{reading.ID}</td>
                <td>{reading.Temperature}</td>
                <td>{reading.Humidity}</td>
                <td>{reading.Timestamp}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>Loading...</p>
      )}
    </div>
  );
};

export default Index;
