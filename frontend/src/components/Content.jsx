import React from 'react'
import Button from './Button'

const Content = ({payload, received}) => {
  return (
    
    <div className='container'>
        <div className='row'>
            <div className='col'>
                <h4 className='mt-5'>Kirim</h4>
                <div className='mt-1' style={{outline: '1px solid #ccc', padding: "2em"}}>
                    <pre>
                        <span className='text-muted'>{payload ? payload : "Belum mengirim apapun"}</span>
                    </pre>
                </div>
            </div>
            <div className='col'>
                <h4 className='mt-5'>Terima</h4>
                <div className='mt-1' style={{outline: '1px solid #ccc', padding: "2em"}}>
                    <pre >
                        <span className='text-muted'>{received ? received : "Belum menerima apapun"}</span>
                    </pre>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Content