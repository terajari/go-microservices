import React from 'react'

const Content = () => {
  return (
    <div className='container'>
        <div className='row'>
            <div className='col'>
                <h4 className='mt-5'>Kirim</h4>
                <div className='mt-1' style={{outline: '1px solid #ccc', padding: "2em"}}>
                    <pre id='payload'>
                        <span className='text-muted'>Belum mengirim apapun...</span>
                    </pre>
                </div>
            </div>
            <div className='col'>
                <h4 className='mt-5'>Terima</h4>
                <div className='mt-1' style={{outline: '1px solid #ccc', padding: "2em"}}>
                    <pre id='received'>
                        <span className='text-muted'>Belum menerima apapun...</span>
                    </pre>
                </div>
            </div>
        </div>
    </div>
  )
}

export default Content