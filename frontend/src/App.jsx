import Header from './components/Header'
import Footer from './components/Footer'
import Content from './components/Content'
import React from 'react';

function App() {
    const [output, setOutput] = React.useState("");
    const [received, setReceived] = React.useState("");
    const [payload, setPayload] = React.useState("");

  return (
    <div>
      <Header output={output} setOutput={setOutput} setReceived={setReceived} setPayload={setPayload} />
      <Content payload={payload} received={received} />
      <Footer />
    </div>
  )
}

export default App
