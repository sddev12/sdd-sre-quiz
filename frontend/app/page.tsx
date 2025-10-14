"use client";
import { useState } from "react";

export default function Home() {
  const [name, setName] = useState("");
  const isNameEntered = name.trim().length > 0;
  return (
    <div className="min-h-screen bg-white text-black font-mono flex flex-col">
      {/* HEADER */}
      <header className="w-full bg-black text-white py-4 px-6 flex items-center justify-center border-b-4 border-black">
        <span className="text-2xl tracking-widest font-bold uppercase">SamOS Terminal</span>
      </header>

      {/* MAIN PANEL */}
      <main className="flex-1 flex flex-col items-center justify-center p-6">
        <section
          className="w-full max-w-xl border-4 border-black p-8 flex flex-col items-center gap-6 bg-white"
        >
          <h1 className="text-2xl md:text-3xl font-bold uppercase tracking-widest text-center mb-2">SRE QUIZ</h1>
          <p className="text-base md:text-lg uppercase tracking-wide text-center mb-4">
            Welcome to the Site Reliability Engineering Quiz<br />
          </p>
          <div className="w-full max-w-xs mb-4 flex flex-col items-center">
            <input
              type="text"
              placeholder="ENTER YOUR NAME"
              className="w-full bg-white text-black text-sm font-mono uppercase tracking-widest outline-none border-0 border-b-4 border-black px-4 py-2 text-center placeholder-black placeholder-opacity-60 focus:ring-0 focus:border-black"
              aria-label="Enter your name"
              value={name}
              onChange={e => setName(e.target.value)}
              style={{ caretColor: '#000', caretShape: 'block' }}
            />
            {/* Custom flashing block cursor for terminal look */}
            <style jsx>{`
              input:focus {
                caret-color: #000;
                caret-shape: block;
              }
              input:focus::placeholder {
                color: transparent;
              }
              input:focus::after {
                content: '';
                display: inline-block;
                width: 10px;
                height: 1.2em;
                background: #000;
                animation: blink 1s steps(1) infinite;
                vertical-align: bottom;
                margin-left: -10px;
                position: relative;
                top: 2px;
              }
              @keyframes blink {
                0%, 50% { opacity: 1; }
                51%, 100% { opacity: 0; }
              }
            `}</style>
          </div>
          <button
            className="uppercase border-4 border-black px-8 py-2 text-sm font-bold tracking-widest bg-white hover:bg-[#EAEAEA] transition-colors duration-100 focus:outline-none focus:ring-2 focus:ring-black disabled:opacity-40 disabled:cursor-not-allowed"
            tabIndex={0}
            disabled={!isNameEntered}
          >
            Begin Quiz
          </button>
        </section>
      </main>

      {/* FOOTER */}
      <footer className="w-full border-t-4 border-black py-2 px-6 flex items-center justify-between text-xs uppercase bg-white">
        <span>VERSION 0.1</span>
        <span>&copy; {new Date().getFullYear()} SRE QUIZ</span>
      </footer>
    </div>
  );
}
