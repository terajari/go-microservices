import React from 'react';
import axios from 'axios';

const Button = ({ name, setOutput, setReceived, setPayload }) => {
  const handleClick = async () => {
    try {
      const response = await axios.post('http://localhost:8080/broker');

      setPayload("Request Post kosong");
      setReceived(JSON.stringify(response.data, null, 4));

      if (response.data.error) {
        setOutput(response.data.error);
      } else {
        setOutput("data dari: " + response.data.message);
      }
    } catch (error) {
      setOutput("Error: " + error);
    }
  };

  return (
    <div>
      <button id='brokerBtn' type="button" className="btn btn-primary" onClick={handleClick}>
        {name}
      </button>
    </div>
  );
};

export default Button;