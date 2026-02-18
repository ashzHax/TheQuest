import { useEffect } from "react";
import { createPortal } from "react-dom";
import style from "./Modal.module.css";

function Modal({ open, message, onClose }) {
  const modalRoot = document.getElementById("modal-root");
  const lines = message.split("\n");

  useEffect(() => {
    const handleEsc = (e) => {
      if (e.key === "Escape") onClose();
    };

    if (open) {
      document.addEventListener("keydown", handleEsc);
      document.body.style.overflow = "hidden"; // prevent scroll
    }

    return () => {
      document.removeEventListener("keydown", handleEsc);
      document.body.style.overflow = "auto";
    };
  }, [open, onClose]);

  if (!open) return null;

  return createPortal(
    <div className={style.overlay} onClick={onClose}>
      <div
        className={style.modal}
        onClick={(e) => e.stopPropagation()} // prevent close when clicking inside
      >
        <div className={style.modalTextBox}>
          {
            lines.map((line, index) => (
              <span className={style.message} key={index}>
                {line}
                {index !== lines.length - 1 && <br />}
              </span>
            ))
          }
        </div>
        <button className={style.button} onClick={onClose}>확인</button>
      </div>
    </div>,
    modalRoot
  );
}

export default Modal;
