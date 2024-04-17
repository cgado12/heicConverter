import { Button, Group, Text, rem } from "@mantine/core";
import { Dropzone } from "@mantine/dropzone";
import { IconPhoto, IconUpload, IconX } from "@tabler/icons-react";
import { FileConverterDialog } from "../../wailsjs/go/main/App";
import { useState } from "react";

import "./Dropzone.css";

const DropzoneArea = () => {
  const [loading, setLoading] = useState(false);

  const convertFilesViaBackendDialog = async () => {
    setLoading(true);
    await FileConverterDialog();
    setLoading(false);
  };

  const DragAndDroppedFiles = async (files: File[]) => {
    setLoading(true);

    const formData = new FormData();
    files.forEach((file) => {
      formData.append("upload", file);
    });

    await fetch("http://localhost:3269/convert", {
      method: "POST",
      body: formData,
    });

    setLoading(false);
  };

  return (
      <Dropzone
        className="dropzone-container"
        onClick={convertFilesViaBackendDialog}
        onDrop={(files) => DragAndDroppedFiles(files)}
        loading={loading}
      >
        <Group justify="center" gap="xl" style={{ pointerEvents: "none" }}>
          <Dropzone.Accept>
            <IconUpload
              style={{
                width: rem(52),
                height: rem(52),
                color: "var(--mantine-color-blue-6)",
              }}
              stroke={1.5}
            />
          </Dropzone.Accept>
          <Dropzone.Reject>
            <IconX
              style={{
                width: rem(52),
                height: rem(52),
                color: "var(--mantine-color-red-6)",
              }}
              stroke={1.5}
            />
          </Dropzone.Reject>
          <Dropzone.Idle>
            <IconPhoto
              style={{
                width: rem(52),
                height: rem(52),
                color: "var(--mantine-color-dimmed)",
              }}
              stroke={1.5}
            />
          </Dropzone.Idle>

          <div>
            <Text size="xl" inline>
              Please Click or Drag and Drop to convert files
            </Text>
            <Text size="sm" c="dimmed" inline mt={7}>
              Attach as many files as you like!
            </Text>
          </div>
        </Group>
      </Dropzone>
  );
};

export default DropzoneArea;
