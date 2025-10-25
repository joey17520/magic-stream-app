import { createContext, useState, useEffect } from "react";

// eslint-disable-next-line react-refresh/only-export-components
export const AuthContext = createContext({});

export default function AuthProvider({ children }) {
  // 从localStorage中初始化auth状态
  const [auth, setAuth] = useState(() => {
    const savedUser = localStorage.getItem("user");
    return savedUser ? JSON.parse(savedUser) : null;
  });

  // 当auth状态变化时，同步更新localStorage
  useEffect(() => {
    if (auth) {
      localStorage.setItem("user", JSON.stringify(auth));
    } else {
      localStorage.removeItem("user");
    }
  }, [auth]);

  return <AuthContext.Provider value={{ auth, setAuth }}>{children}</AuthContext.Provider>;
}
