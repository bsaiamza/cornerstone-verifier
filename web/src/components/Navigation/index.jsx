// react
import { useState } from 'react'
// mui5 icons
import { Menu as MenuIcon } from '@mui/icons-material'
// components
import AppBarComponent from '../AppBar'
import BoxComponent from '../Box'
import IconButtonComponent from '../IconButton'
import ToolbarComponent from '../Toolbar'
import TooltipComponent from '../Tooltip'
// images
import logo from '../../assets/images/logo_dark.png'
// sidebar
import Sidebar from './sidebar'
import { Link } from 'react-router-dom'

const Navigation = () => {
  const [openDrawer, setOpenDrawer] = useState(false)

  const handleOpenDrawer = () => setOpenDrawer(true)
  const handleCloseDrawer = () => setOpenDrawer(false)

  const renderSidebar = (
    <Sidebar handleCloseDrawer={handleCloseDrawer} openDrawer={openDrawer} />
  )

  return (
    <>
      <AppBarComponent position="sticky" sx={{ marginBottom: '1rem' }}>
        <ToolbarComponent>
          <>
            <Link to="/">
              <BoxComponent
                alt="DI"
                component="img"
                src={logo}
                sx={{
                  height: '5rem',
                  width: '5rem',
                }}
              />
            </Link>
            <BoxComponent sx={{ flexGrow: 1 }} />
            <IconButtonComponent>
              <IconButtonComponent
                ariaLabel="menu"
                color="white"
                edge="end"
                onClick={handleOpenDrawer}
                size="large"
              >
                <TooltipComponent title="Menu">
                  <MenuIcon />
                </TooltipComponent>
              </IconButtonComponent>
            </IconButtonComponent>
          </>
        </ToolbarComponent>
      </AppBarComponent>
      {renderSidebar}
    </>
  )
}

export default Navigation
