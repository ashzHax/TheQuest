import { useEffect } from "react";
import { createPortal } from "react-dom";
import style from "./Modal.module.css";

function Modal({ open, message, onClose }) {
  const modalRoot = document.getElementById("modal-root");

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
        <p className={style.message}>{message}</p>
        <button className={style.button} onClick={onClose}>OK</button>
      </div>
    </div>,
    modalRoot
  );
}

export default Modal;
