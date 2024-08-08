
const configPOST = (data) => {
   return {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      referrerPolicy: "no-referrer",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify(data),
    }
}

const configGET = () => {
   return {
      method: "GET",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      referrerPolicy: "no-referrer",
      headers: {
        "Content-Type": "application/json"
      }
    }
}

/**
 * configPOSTAccessData
 * @param {*} data 
 * @param {*} accessData 
 * @returns 
 */
const configPOSTAccessData = (data, accessData) => {
   return {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      referrerPolicy: "no-referrer",
      headers: {
        "Content-Type": "application/json",
        "serial_user": accessData.serial_user,
        "serial_company": accessData.serial_company,
        "email": accessData.email,
        "u": accessData.u
      },
      body: JSON.stringify(data),
    }
}

/**
 * configPOSTAccessData
 * @param {*} data 
 * @param {*} accessData 
 * @returns 
 */
const configPOSTUpload = (data, accessData, type) => {
   return {
      method: "POST",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      referrerPolicy: "no-referrer",
      headers: {
        "Content-Type": "multipart/form-data",
        "serial_user": accessData.serial_user,
        "serial_company": accessData.serial_company,
        "email": accessData.email,
        "type": type,
        "u": accessData.u
      },
      body: JSON.stringify(data),
    }
}

/**
 * configGETAccessData
 * @param {*} accessData 
 * @returns 
 */
const configGETAccessData = (accessData) => {
   return {
      method: "GET",
      mode: "cors",
      cache: "no-cache",
      credentials: "same-origin",
      referrerPolicy: "no-referrer",
      headers: {
        "Content-Type": "application/json",
        "serial_user": accessData.serial_user,
        "serial_company": accessData.serial_company,
        "email": accessData.email,
        "u": accessData.u
      }
    }
}

/**
 * translateToResponseEntity
 * @param {*} fetchResponse 
 * @returns 
 */
const translateToResponseEntity = async (fetchResponse) => {
   let responseEntity = new ResponseEntity();
   const rawResponse = await fetchResponse.json();
   responseEntity = rawResponse;
   return responseEntity;
}

/**
 * callPOST
 * @param {*} url 
 * @param {*} data 
 * @returns 
 */
const callPOSTAccessData = async (url, data, accessData) => {
   const result = await fetch(url,configPOSTAccessData(data,accessData));
   const responseEntity = await translateToResponseEntity(result);
   return responseEntity;
}

/**
 * callGET
 * @param {*} url 
 * @returns 
 */
const callGET = async (url) => {
   const result = await fetch(url,configGET());
   const responseEntity = await translateToResponseEntity(result);
   return responseEntity;
}

/**
 * callGET
 * @param {*} url 
 * @returns 
 */
const callPOST = async (url, data) => {
   const result = await fetch(url,configPOST(data));
   const responseEntity = await translateToResponseEntity(result);
   return responseEntity;
}

/**
 * callGETAccessData
 * @param {*} url 
 * @returns 
 */
const callGETAccessData = async (url,accessData) => {
   const result = await fetch(url,configGETAccessData(accessData));
   const responseEntity = await translateToResponseEntity(result);
   return responseEntity;
}