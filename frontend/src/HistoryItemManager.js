const actionName = {
    0: "New stonecutter recipe",
    1: "New furnace recipe",
    2: "New furnace and smoker recipe",
    3: "New regular and blast furnace recipe",
    4: "New crafting table recipe",
    5: "Recipe removed"
};
const actionCategory = {
    0: "New recipe",
    1: "Recipe removed"
};
export function actionTypeToString(actionType) {
    return actionName[actionType] || "Unknown ActionType";
}
function IdentifyActionCategory(actionTypeKey) {
    if (actionTypeKey >= 0 && actionTypeKey <= 4) {
        return actionCategory[0];
    } else if (actionTypeKey === 5) {
        return actionCategory[1];
    } else {
        return null;
    }
}
export function historyItemLabel(historyItem) {
    console.log(historyItem);
    const category=IdentifyActionCategory(historyItem.ActionType);
    if (category === "New recipe") {
        if (historyItem.Arguments.hasOwnProperty('input') && historyItem.Arguments.hasOwnProperty('output')) {
            return `'${historyItem.Arguments['input']}' to '${historyItem.Arguments['output']}'`;
        }
        if (historyItem.ActionType === 4) {
            //for crafting table
            return `'${historyItem.Arguments['output']}' ` + (historyItem.Arguments['isShapeless'] === 'true' ? 'shapeless' : 'shaped') + ' crafting recipe';
        }
    }
    else if (category==="Recipe removed")
    {
        return `removed '${historyItem.Arguments['item']}' creation recipe`;
    }
    return '';
}