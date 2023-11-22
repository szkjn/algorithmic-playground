/*
Code Wars Kata :

Your task is to implement a function that would take the morse code as 
input and return a decoded human-readable string.
*/

const MORSE_CODE = {
    '.-': 'A', '-...': 'B', '-.-.': 'C', '-..': 'D', '.': 'E', '..-.': 'F',
    '--.': 'G', '....': 'H', '..': 'I', '.---': 'J', '-.-': 'K', '.-..': 'L',
    '--': 'M', '-.': 'N', '---': 'O', '.--.': 'P', '--.-': 'Q', '.-.': 'R',
    '...': 'S', '-': 'T', '..-': 'U', '...-': 'V', '.--': 'W', '-..-': 'X',
    '-.--': 'Y', '--..': 'Z', '-----': '0', '.----': '1', '..---': '2', 
    '...--': '3', '....-': '4', '.....': '5', '-....': '6', '--...': '7', 
    '---..': '8', '----.': '9', '.-.-.-': '.', '--..--': ',', '..--..': '?', 
    '.----.': "'", '-.-.--': '!', '-..-.': '/', '-.--.': '(', '-.--.-': ')', 
    '.-...': '&', '---...': ':', '-.-.-.': ';', '-...-': '=', '.-.-.': '+', 
    '-....-': '-', '..--.-': '_', '.-..-.': '"', '...-..-': '$', '.--.-.': '@', 
    '...---...': 'SOS'
};

function decodeMorse(morseCode) {
    const words = morseCode.trim().split('   ');
    const result = words.map(word => 
        word.split(' ').map(letter => MORSE_CODE[letter]).join('')
    );

    return result.join(' ');
}

console.log(decodeMorse('.... . -.--   .--- ..- -.. .'));  // Output: 'HEY JUDE'
console.log(decodeMorse('...---...'));                      // Output: 'SOS'
console.log(decodeMorse('...   ---   ...'));                // Output: 'S O S'
