import { Alert, Group, Text, rem } from "@mantine/core";
import { Dropzone } from "@mantine/dropzone";
import { IconPhoto, IconUpload, IconX } from "@tabler/icons-react";
import { FileConverter } from "../../wailsjs/go/main/App";
import "./Dropzone.css";
import { useState } from "react";

const DropzoneArea = () => {
  const [loading, setLoading] = useState(false)

  const convertFiles = async () => {
    setLoading(true)
    await FileConverter();
    setLoading(false)
  };

  const testFileConv = (f: any) => {
    console.log(f)
  }
  console.log("unsupported")
  return (
    <Dropzone
      className="dropzone-container"
      onClick={convertFiles}
      onDrop={(files)=> testFileConv(files)}
      loading={loading}
      //   maxSize={5 * 1024 ** 2}
      // accept={MIME_TYPE}
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
            Please Click to select files
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
