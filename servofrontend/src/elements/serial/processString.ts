
import { getProgramEnd, getProgramStart, postToolToQueue, postToolLength} from "../../api/api"

export const processString = (s:string)=>{

    if(s.includes("M30")){
        console.log("end of program detected")
        getProgramEnd()
    } else if (s.includes("O0001")){
        console.log("start of program detected")
        getProgramStart()
    }else if (s.includes("M06")){
        console.log("toolchange detected")
        const toolNumber = getNumberToolChange(s)
        if(toolNumber<0){
            console.error("error finding tools in command string",s)
        }
        postToolToQueue({
            toolid:toolNumber
        })
        console.log("posted tool", toolNumber)
    }else if(s.includes("TO")){
        console.log("tool offset detected", s)

       const { toolNumber, offset} =   getToolOffsetNumbers(s)

       if(toolNumber == null || offset ==null){
        console.error("unable to detect tool length")
        return
       }
       postToolLength({
        toolid:toolNumber,
        toollength: offset
       })
       console.log("posted length", {
        toolid:toolNumber,
        toollength: offset
       })

    }
}


function getNumberToolChange(str:string): number {
    // Use regular expression to match the numbers after "T"
    const matches = str.match(/T(\d{2})/);

    // If there are matches and it has captured groups
    if (matches && matches[1] ) {
        // Return the captured numbers as an array of two integers
        return parseInt(matches[1])
    } else {
        // If no matches found or no captured groups, return null or appropriate value
        return -1;
    }
}

function getToolOffsetNumbers(inputString:string) {
    // Split the string by spaces
    const parts = inputString.split(' ');

    let toolNumber = null;
    let offset = null;

    console.log("parts", parts    )
    // Loop through the parts of the string
    for (let i = 0; i < parts.length; i++) {
        const part = parts[i];
        console.log("checking part", part, "i", i)
        // Check if the current part is 'TO'
        if (part === 'TO' && i + 1 < parts.length) {
            // Get the number after 'TO'
            console.log("part1", parts[i + 1])
            toolNumber = parseInt(parts[i + 1]);
        }
        // Check if the current part is a number
        else if (!isNaN(parseFloat(part)) && i+1 == parts.length) {
            // Update the last number found
            offset = parseFloat(part);
        }
    }

    console.log({toolNumber, offset})
    return {
        toolNumber,
        offset
    };
}