/**
 * Track the trade of a commodity from one trader to another
 * @param {org.example.mynetwork.Trade} trade - the trade to be processed
 * 
 
async function tradeCommodity(trade) {

    // set the new owner of the commodity
    trade.commodity.owner = trade.newOwner;
    let assetRegistry = await getAssetRegistry('org.example.mynetwork.Commodity');

    // emit a notification that a trade has occurred
    let tradeNotification = getFactory().newEvent('org.example.mynetwork', 'TradeNotification');
    tradeNotification.commodity = trade.commodity;
    emit(tradeNotification);

    // persist the state of the commodity
    await assetRegistry.update(trade.commodity);
}
*/
/**
 * Remove all high volume commodities
 * @param {org.example.mynetwork.RemoveHighQuantityCommodities} remove - the remove to be processed
 * 
 
async function removeHighQuantityCommodities(remove) {

    let assetRegistry = await getAssetRegistry('org.example.mynetwork.Commodity');
    let results = await query('selectCommoditiesWithHighQuantity');

    for (let n = 0; n < results.length; n++) {
        let trade = results[n];

        // emit a notification that a trade was removed
        let removeNotification = getFactory().newEvent('org.example.mynetwork','RemoveNotification');
        removeNotification.commodity = trade;
        emit(removeNotification);
        await assetRegistry.remove(trade); }
}
*/

/**
 * Track the trade of a commodity from one trader to another
 * @param {org.example.mynetwork.Sale} sale - the trade to be processed
 * @transaction
 */ 
async function saleEmployee(sale) {

    // set the new owner of the commodity
    sale.employee.owner = sale.comapanyname;
    let assetRegistry = await getAssetRegistry('org.example.mynetwork.Employee');
  
    // persist the state of the commodity
    await assetRegistry.update(sale.employee);
}



/**
 * Select all hitachi employees
 * @param {org.example.mynetwork.SelectHitachiEmployees} select - select hitachi employees
 * @transaction
 */

async function selectHitachiEmployees(){
   let assetRegistry = await getAssetRegistry('org.example.mynetwork.Employee');
   let results = await query('selectHitachiEmployees');   
}

/**
 * Select all hitachi employees
 * @param {org.example.mynetwork.SelectHumanEmployees} select - select hitachi employees
 * @transaction
 */

async function selectHitachiEmployees(){
   let assetRegistry = await getAssetRegistry('org.example.mynetwork.Employee');
   let results = await query('selectHumanEmployees');   
}
