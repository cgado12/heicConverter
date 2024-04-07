import React from 'react'
import {createRoot} from 'react-dom/client'
import './style.css'
import App from './App'
import { ColorSchemeScript, createTheme, MantineProvider } from '@mantine/core';


const container = document.getElementById('root')

const root = createRoot(container!)

// const theme = createTheme({
//     /** Put your mantine theme override here */
//   });

root.render(
    <React.StrictMode>
        <MantineProvider defaultColorScheme='dark'>
        <ColorSchemeScript />
        <App/>
        </MantineProvider>
    </React.StrictMode>
)
