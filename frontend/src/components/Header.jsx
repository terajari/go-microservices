import React from 'react'

const Header = () => {
    
  return (
    <div className='container text-center'>
        <div className='row'>
            <div className='col'>
                <h1>Go Microservices</h1>
                <hr />
                <div id='output' className='mt-5' style={{outline: '1px solid #ccc', padding: "2em"}}>
                    <span className='text-muted'>
                        Hasil di sini...
                    </span>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Header