export function isJSON(jsonString = '') {
    try {
        return JSON.parse(jsonString);
    } catch (e) {
        return false;
    }
}
