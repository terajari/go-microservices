import React from 'react'
import Button from './Button'

const Header = ({output, setOutput, setReceived, setPayload}) => {
    
  return (
    <div className='container text-center'>
        <div className='row'>
            <div className='col'>
                <h1>Go Microservices</h1>
                <hr />
                <Button name='Broker Service' setOutput={setOutput} setReceived={setReceived} setPayload={setPayload} />
                <div className='mt-5' style={{outline: '1px solid #ccc', padding: "2em"}}>
                    <span className='text-muted'>
                        {output ? output : "Hasil di sini..."}
                    </span>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Header