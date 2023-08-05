//paste in browser developer tools console after game loaded
let output = '';
EJS_emulator.Module.cwrap('get_core_options', 'string', [])().split('\n').forEach(line => {
    const parts = line.split(';');
    const nameParts = parts[0].split('|');
    output += `{Id:"${nameParts[0]}",Name:"${nameParts[0]}",Variants:"${parts[1].trim()}",Default:"${nameParts[1]}",},`;
});
output
