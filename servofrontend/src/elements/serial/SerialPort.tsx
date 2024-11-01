import { Button } from "react-bootstrap";
import { processString } from "./processString";

export const SerialPort = (): React.JSX.Element => {


  const openSerialPort = async () => {
    if (!("serial" in navigator)) {
      return;
    }
    const port = await navigator.serial.requestPort();
    await port.open({
      baudRate: 6000,
    });

    while (port.readable) {
      const reader = port.readable.getReader();

      try {
        let readString = ""

        // eslint-disable-next-line no-constant-condition
        while (true) { 
          const { value, done } = await reader.read();
          if (done) {
            // Allow the serial port to be closed later .
            reader.releaseLock();
            break;
          }
          if (value) {

            const string = new TextDecoder().decode(value);
            
            for(const c of string){
              if(c == "\n"){
                processString(readString)
                readString = ""
              } else if (c =='\r' || c =='\x00' ){
                //BALEET
              } else{
                readString += c
              }
            }
            
     
            
          }
        }
      } catch (error) {
        console.error("serial error", error);

      }
    }
  };
  return <Button onClick={() => openSerialPort()}> Open Serial Port </Button>;
};
