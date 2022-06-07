// components
import Copyright from './Copyright'

const Footer = () => {
  return (
    <div
      style={{
        position: 'fixed',
        bottom: 0,
        right: 0,
        backgroundColor: '#eee',
        padding: '2rem',
        marginTop: '5rem',
        zIndex: '1',
      }}
    >
      <Copyright />
    </div>
  )
}

export default Footer
