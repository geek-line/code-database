import React from 'react'
import { createRoot } from 'react-dom/client'
import TagRanking from '../../components/TagRanking'
import 'materialize-css/dist/css/materialize.min.css'
import './style.css'

const containers = document.querySelectorAll('#tag_ranking')
containers.forEach((container) => {
  if (container) {
    const root = createRoot(container)
    root.render(<TagRanking />)
  }
})
