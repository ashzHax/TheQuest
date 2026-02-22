import './font/pretendard.css'

import { useState } from 'react';
import Modal from "./Modal.jsx";
import style from './App.module.css'

export default function App() {
  const [letters, setLetters] = useState(["", "", "", "", "", "", "", "", "", ""]);

  const [input, setInput] = useState("");
  const [modalMessage, setModalMessage] = useState("");

  const handleSubmit = async (e) => {
    e.preventDefault();

    const textToSend = input;   // save before clearing
    setInput("");               // clear input immediately

    try {
      const res = await fetch("/api", {
        method: "POST",
        headers: {
          "Content-Type": "application/json"
        },
        body: JSON.stringify({ text: textToSend })
      });

      const data = await res.json();

      switch (data.action) {
        case "popup":
          setModalMessage(data.message);
          break;

        case "fill":
          setLetters(prev => {
            const updated = [...prev]; // clone (important)

            data.fills.forEach(fill => {
              updated[fill.location] = fill.letter;
            });

            return updated;
          });
          break;

        default:
          console.warn("Unknown action:", data);
      }
    } catch (err) {
      console.error(err);
    }
  };

  return (
      <div className={style.App}>

        <div className={style.wordleBox}>
          {letters.map(letter => (
            <div className={style.blank}>
              <p className={style.blankWord}>{letter}</p>
            </div>
          ))}
        </div>

        <form onSubmit={handleSubmit} className={style.inputForm}>
          <input 
            className={style.inputBox} 
            value={input}
            onChange={(e) => setInput(e.target.value)}
            type="text" placeholder="정답을 입력하세요" 
          />
          <button className={style.inputSubmitButton} type="submit">제출</button>
        </form>

        <Modal
          open={!!modalMessage}
          message={modalMessage}
          onClose={() => setModalMessage("")}
        />
      </div>
    )
}