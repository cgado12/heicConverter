import { useEffect, useState } from "react";
import { ActionIcon, Button, Container, Text } from "@mantine/core";
import "@mantine/core/styles.css";
import "@mantine/dropzone/styles.css";
import { GetOutDir, Greet, SetOutDir } from "../wailsjs/go/main/App";
import DropzoneArea from "./components/Dropzone";
import { IconFolder } from "@tabler/icons-react";

import "./App.css";

const App = () => {
  const [outDir, setDir] = useState<string | undefined>(undefined);

  useEffect(() => {
    const getDir = async () => {
      const dir = await GetOutDir();
      setDir(dir);
    };
    if (!outDir) {
      getDir();
    }
  }, []);

  const setOutDirectory = async () => {
    const dir = await SetOutDir();
    setDir(dir)
  };

  return (
    <div id="App">
      <div className="output-dir-container">
        <Text>Output Directory: {outDir}</Text>
        <ActionIcon size="md" variant="outline" onClick={setOutDirectory}>
          <IconFolder size={32} />
        </ActionIcon>
      </div>
      <div className="container">
        <DropzoneArea />
      </div>
    </div>
  );
};

export default App;
