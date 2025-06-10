import '../Global.css'
import logo from '../../assets/logo.png'
import Form from 'react'

function Header(){
  return (
    <header className="bg-white shadow-md sticky top-0 z-50">
      <nav className="container mx-auto px-6 py-3 flex space-x-5 justify-between items-center">
        <div className="flex items-center">
          <img src={logo} alt="Company Logo" className="h-8" />
        </div>

        <ul className="space-x-4 flex   w-full">
          <a href="">
            <li>Home</li>
          </a>
          <a href="">
            <li>Services</li>
          </a>
          <a href="">
            <li>Portfolio</li>
          </a>
          <a href="">
            <li>Testimonials</li>
          </a>
          <a href="">
            <li>Team</li>
          </a>
          <a href="">
            <li>
              <div className="flex-col w-16">
                <span>Menu</span>
                <img className="w-3 bg-amber-50" src="./down.svg" alt="" />
                <div id="menu" className="hidden">
                  <ul>
                    <a href="">
                      <li>About</li>
                    </a>
                    <a href="">
                      <li></li>
                    </a>
                    <a href="">
                      <li></li>
                    </a>
                    <a href="">
                      <li></li>
                    </a>
                    <a href="">
                      <li></li>
                    </a>
                  </ul>
                </div>
              </div>
            </li>
          </a>
          <a href="">
            <li>News</li>
          </a>
          <a href="">
            <li>
              <button className="bg-black-300 b">Get Quotes</button>
            </li>
          </a>
          <a href="">
            <li>
              <button className="bg-blue-300">Theme</button>
            </li>
          </a>
        </ul>
      </nav>

      <div
        id="mobile-menu"
        className="hidden md:hidden bg-white px-6 py-3 border-t"
      >
        <a href="#" className="block py-2 text-gray-700 hover:text-blue-600">
          Home
        </a>
        <a href="#" className="block py-2 text-gray-700 hover:text-blue-600">
          Services
        </a>
        <a href="#" className="block py-2 text-gray-700 hover:text-blue-600">
          Portfolio
        </a>
        <a href="#" className="block py-2 text-gray-700 hover:text-blue-600">
          Testimonials
        </a>
        <a href="#" className="block py-2 text-gray-700 hover:text-blue-600">
          Team
        </a>
        <a href="#" className="block py-2 text-gray-700 hover:text-blue-600">
          News
        </a>
        <div className="py-2">
          <button className="bg-gray-100 hover:bg-gray-200 px-4 py-2 rounded-lg font-medium w-full">
            Get Quotes
          </button>
        </div>
      </div>
    </header>
  );}



export default Header;
