import axios from "./apiConfig";

export const getReadings = async () => {
  try {
    const response = await axios.get(`/readings`);
    return response.data;
  } catch (err) {
    console.log("Couldn't get readings:", err);
    throw err;
  }
};
