import React from 'react';
import axios from 'axios';

const Button = ({ name, url, setOutput, setReceived, setPayload }) => {

  let handleClick = () => {};

  switch (name) {
    case "Broker Service":
      handleClick = async () => {
        try {
          const response = await axios.post(url);
    
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
      break;
    
    case "Auth Service":
      handleClick = async () => {
        const payload = {
          action: "auth",
          auth: {
            email: "admin@example.com",
            password: "verysecret"
          }
        }

        try {
          const response = await axios.post(url, payload);

          setPayload(JSON.stringify(payload, null, 4));
          setReceived(JSON.stringify(response.data, null, 4));

          if (response.data.error) {
            setOutput(response.data.error);
          } else {
            setOutput("data dari: " + response.data.message);
          }
        } catch(error) {
          setOutput("Error: " + error);
        }
      }
      break;

      default:
        break;
  }
  return (
    <div>
      <button id='brokerBtn' type="button" className="btn btn-primary" onClick={handleClick}>
        {name}
      </button>
    </div>
  );
};

export default Button;