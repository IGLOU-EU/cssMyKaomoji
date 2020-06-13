/* *
 * MIT License
 * 
 * Copyright (c) 2020  Gino Ladowitch
 * 
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 * 
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 * 
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 * */

const emojis = Array.from(document.getElementsByClassName('kaomoji'));

const copyToClipBoard = (text) => {
    navigator.clipboard.writeText(text);
    console.log('copied');
}

const infoClipBoard = (text) => {
    info.classList.add('active');
    info.textContent = `'${text}' Copied to clipboard`;
    setTimeout( () => {
        info.classList.remove('active');
    }, 2000);
}

emojis.forEach(emoji => {

    emoji.addEventListener(
        'click', (event) => {
        let contentBefore = getComputedStyle(emoji, ':before').getPropertyValue('content').replace('"', '').replace('"', '');
        copyToClipBoard(contentBefore);
        infoClipBoard(contentBefore);
        }
    );

    emoji.getElementsByTagName('sup')[0].addEventListener(
        'click', (event) => {
            event.stopPropagation();
            let className = event.target.parentElement.classList[1];
            copyToClipBoard(className);
            infoClipBoard(className);
        }
    );

    const info = document.getElementById("info");
});