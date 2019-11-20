/**
@apiDefine  CreateUserError

**/


/**
 * @api {post} /api/user/create 建立用戶
 * @apiVersion 0.0.1
 * @apiName PostUser
 * @apiGroup User
 *
 * @apiDescription 建立新用戶
 *
 * @apiParam {String} name 用戶姓名
 * @apiParam {String} email 用戶Email
 * @apiParam {String} company 用戶公司名稱
 * @apiParam {String} mobile  用戶手機號碼
 *
 * @apiSuccess {Number} code  錯誤代碼
 *                 0:SUCCESS(成功)
 *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
 *                 2:USER_EXIST(用戶(Email)已存在)
 * @apiSuccess {String} message  錯誤訊息
 * @apiSuccess {String} id   用戶ID (若是成功建立)
 *
 */
 function postUser() { return; }

 /**
  * @api {post} /api/user/info  尋找用戶
  * @apiVersion 0.0.1
  * @apiName GetUser
  * @apiGroup User
  *
  * @apiDescription 使用用戶的ID,Email或Mobile其中之一來取得用戶資訊，參數三種至少要有一種
  *
  * @apiParam {String} id  用戶 ID
  * @apiParam {String} email 用戶 Email
  * @apiParam {String} mobile 用戶 Mobile
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
  *                 3:USER_NOT_EXIST(用戶不存在)
  * @apiSuccess {String} message  錯誤訊息
  * @apiSuccess  {Object} user         用戶資訊
  * @apiSuccess  {String} user.id
  * @apiSuccess  {String} user.name
  * @apiSuccess  {String} user.email
  * @apiSuccess   {String}  user.company
  * @apiSuccess   {String}  user.mobile
  * @apiSuccess   {Boolean}  user.registered 是否已報到
  * @apiSuccess   {Boolean}  user.frEnabled 是否已註冊照片
  * @apiSuccess   {String} user.imageUrl  用戶圖片url
  * @apiSuccess   {Object[]} user.checkList  用戶到攤位簽到紀錄
  * @apiSuccess   {String} user.checkList.boothName  簽到紀錄-攤位名稱
  * @apiSuccess   {Boolean} user.checkList.checked  簽到紀錄-是否已簽到
  * @apiSuccess   {Number}  user.checkList.time 簽到紀錄-簽到時間unix time
  *
  */

 function getUser() { return; }


 /**
  * @api {post} /api/user/check  攤位簽到
  * @apiVersion 0.0.1
  * @apiName UpdateUserCheck
  * @apiGroup User
  *
  * @apiDescription 若人臉辨識成功，在攤位簽到
  *
  * @apiParam {String} id  用戶 ID
  * @apiParam {String} boothName 攤位名稱
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *
  */
  function updateUserCheck() { return; }

 /**
  * @api {post} /api/user/image  更新用戶圖片
  * @apiVersion 0.0.1
  * @apiName UpdateUserImage
  * @apiGroup User
  *
  * @apiDescription 更新用戶圖片，若格式不符會回傳錯誤 4:INVALID_IMAGE(圖片格式不符)
  *
  * @apiParam {String} id  用戶 ID
  * @apiParam {String} image 用戶臉部圖片(base64-encoded)
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *                 4:INVALID_IMAGE(圖片格式不符)
  *
  */
function updateUserImage() { return; }

  /**
   * @api {post} /api/user/update  更新用戶資訊
   * @apiVersion 0.0.1
   * @apiName UpdateUser
   * @apiGroup User
   *
   * @apiDescription 更新用戶資訊，包含基本資訊及用戶簽到與否
   *
   * @apiParam {String} id  用戶 ID
   * @apiParam {String} name 用戶姓名
   * @apiParam {String} email 用戶Email
   * @apiParam {String} company 用戶公司名稱
   * @apiParam {String} mobile  用戶手機號碼
   * @apiParam {String} registered 用戶是否已經簽到
   *
   * @apiSuccess {Number} code  錯誤代碼
   *                 0:SUCCESS(成功)
   *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
   *                 3:USER_NOT_EXIST(用戶不存在)
   *
   */
  function updateUser() { return; }

  /**
   * @api {post} /api/user/delete  刪除用戶
   * @apiVersion 0.0.1
   * @apiName DeleteUser
   * @apiGroup User
   *
   * @apiDescription 刪除用戶資料
   *
   * @apiParam {String} id  用戶 ID
   *
   * @apiSuccess {Number} code  錯誤代碼
   *                 0:SUCCESS(成功)
   *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
   *                 3:USER_NOT_EXIST(用戶不存在)
   *
   */

  function deleteUser() { return; }

 function updateUserImage()) { return; }
 /**
  * @api {post} /api/fr/verification 人臉辨識
  * @apiVersion 0.0.1
  * @apiName faceRecognition
  * @apiGroup Face Recognition
  * @apiDescription 人臉辦識，可以選擇可信度，及最大回傳數目，結果會回傳0-max個辨識結果高於可信度值的用戶資訊
  * @apiParam {String} image 臉部圖片(base64-encoded)
  * @apiParam {Number}  threshold  可信度(0.0-1.0),數字愈高可信度愈高
  * @apiParam {Numnber} max  最多回傳用戶數目
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *                 4:INVALID_IMAGE(圖片格式不符)
  * @apiSuccess {Object[]} userList  回傳用戶資訊陣列
  * @apiSuccess   {String}  userList.id
  * @apiSuccess   {String}  userList.email
  * @apiSuccess  {String}   userList.name
  * @apiSuccess  {String}   userList.mobile
  * @apiSuccess   {String}  userList.imageUrl  用戶圖片url
  * @apiSuccess   {Boolean} userList.registered 是否已報到
  */
  function faceRecognition() { return; }
