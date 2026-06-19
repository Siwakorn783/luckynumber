import React from 'react'
import link from "next/link"
import Link from 'next/link'
function Navbar() {
    return (
        <nav className='flex gap-4 border-b p-4'>
            <Link href="/">Randomoneluckynumber</Link>
            <Link href="/luckynumuserdefined">Luckynumuserdefined</Link>

        </nav>
    )
}

export default Navbar