export function actionTypeToString(actionType) {
    const actionTypes = {
        0: "New stonecutter recipe",
        1: "New furnace recipe",
        2: "New furnace and smoker recipe",
        3: "New regular and blast furnace recipe",
        4: "New crafting table recipe",
        5: "Recipe removed"
    };

    return actionTypes[actionType] || "Unknown ActionType";
}
export function historyItemLabel(historyItem) {
    return historyItem.Arguments.output;
}
