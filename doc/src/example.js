/**
@apiDefine  CreateUserError

**/


/**
 * @api {post} /api/user/create 1.建立用戶
 * @apiVersion 0.0.1
 * @apiName PostUser
 * @apiGroup User
 *
 * @apiDescription 建立新用戶, company+firstname+lastname是唯一
 *
 * @apiParam {String}   firstname    用戶名(必填)
 * @apiParam {String}   lastname     用戶姓(必填)
 * @apiParam {String}   email       用戶email(必填)
 * @apiParam {String}   company     用戶公司名稱
 * @apiParam {String}   title       用戶公司職稱
 * @apiParam {String}   mobile      用戶手機號碼
 ** @apiParam {String}  extend1     自訂資料1
 * @apiParam  {String}  extend2     自訂資料2

 * @apiSuccess {Number} code  錯誤代碼
 *                 0:SUCCESS(成功)
 *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
 *                 2:USER_EXIST(用戶已存在)
 *                 5:OPERATION_FAIL(建立失敗)
 * @apiSuccess {String} message  錯誤訊息
 * @apiSuccess {String} id   用戶ID (若是成功建立)
 *
 */
 function postUser() { return; }

 /**
  * @api {post} /api/user/info  6.尋找用戶
  * @apiVersion 0.0.1
  * @apiName GetUser
  * @apiGroup User
  *
  * @apiDescription 使用用戶的ID來取得用戶資訊
  *
  * @apiParam {String} id  用戶 ID
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
  *                 3:USER_NOT_EXIST(用戶不存在)
  * @apiSuccess {String}   message  錯誤訊息
  * @apiSuccess  {Object}  user         用戶資訊
  * @apiSuccess  {String}    user.id    用戶 ID
  * @apiSuccess {String}   firstname    用戶名
  * @apiSuccess {String}   lastname     用戶姓
  * @apiSuccess {String}   email       用戶email
  * @apiSuccess {String}   company     用戶公司名稱
  * @apiSuccess  {String}   mobile      用戶手機號碼
  * @apiSuccess   {Boolean}  user.registered 是否已報到
  * @apiSuccess   {Boolean}  user.counterRegistered 是否由櫃台報到
  * @apiSuccess   {Boolean}  user.registerTime    報到時間
  * @apiSuccess   {Boolean}  user.faceRegistered 是否已註冊照片
  * @apiSuccess   {Object[]} user.checkList     用戶到攤位簽到紀錄
  * @apiSuccess   {String}   user.extend1     自訂資料1
  * @apiSuccess    {String}  user.extend2     自訂資料2
  * @apiSuccess   {String}   user.checkList.boothName  簽到紀錄-攤位名稱
  * @apiSuccess   {Boolean}  user.checkList.checked  簽到紀錄-是否已簽到
  * @apiSuccess   {Number}   user.checkList.time 簽到紀錄-簽到時間unix time
  *
  */

 function getUser() { return; }
 /**
  * @api {post} /api/user/face  5.取得用戶圖片
  * @apiVersion 0.0.1
  * @apiName GetUserFace
  * @apiGroup User
  *
  * @apiDescription 使用用戶的ID來取得用戶圖片
  *
  * @apiParam {String} id  用戶 ID
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
  *                 3:USER_NOT_EXIST(用戶不存在)
  * @apiSuccess {String}   message  錯誤訊息
  * @apiSuccess  {String}  image         用戶圖片(base64)
  *
  */

 function getUserFace() { return; }

 /**
  * @api {post} /api/user/check  4.攤位簽到
  * @apiVersion 0.0.1
  * @apiName UpdateUserCheck
  * @apiGroup User
  *
  * @apiDescription 若人臉辨識成功，在攤位簽到
  *
  * @apiParam {String}     id        用戶 ID
  * @apiParam {String}     boothName 攤位名稱
  * @apiParam {Boolean}    checked  是否已簽到
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *
  */
  function updateUserCheck() { return; }

 /**
  * @api {post} /api/user/updateImage  3.更新用戶圖片
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
function updateUserCheck() { return; }

  /**
   * @api {post} /api/user/update  2.更新用戶資訊
   * @apiVersion 0.0.1
   * @apiName UpdateUser
   * @apiGroup User
   *
   * @apiDescription 更新用戶資訊，包含基本資訊及用戶簽到與否
   *
   * @apiParam {String} id  用戶 ID (必填)
   * @apiParam {String} email 用戶email
   * @apiParam {String} mobile  用戶手機號碼
   * @apiParam {String} title  用戶職稱
   * @apiParam {String}  extend1     自訂資料1
   * @apiParam  {String}  extend2     自訂資料2
   * @apiParam {Bool}     registered 用戶是否已經簽到
   * @apiParam {Bool}     counterRegistered 用戶是否由櫃台簽到
   *
   * @apiSuccess {Number} code  錯誤代碼
   *                 0:SUCCESS(成功)
   *                 1:INVALID_PARAMETERS(參數缺少或錯誤)
   *                 3:USER_NOT_EXIST(用戶不存在)
   *
   */
  function updateUser() { return; }

  /**
   * @api {post} /api/user/delete  8.刪除用戶
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

  /**
   * @api {post} /api/user/list 7.列出使用者
   * @apiVersion 0.0.1
   * @apiName listUsers
   * @apiGroup User
   * @apiDescription 可以透過關鍵字搜索名字/公司名稱及Email，列出含有字串的使用者,不使用關鍵字則列出全部使用者
   * @apiParam {String} keyword 關鍵字
   *
   * @apiSuccess {Number} code  錯誤代不使用關鍵字則列出全部使用者
   *                 0:SUCCESS(成功)
   * @apiSuccess {Object[]} userList  回傳用戶資訊陣列(物件參考/api/user/info)
   */
   function listUers() { return; }
 /**
  * @api {post} /api/fr/verification 人臉辨識
  * @apiVersion 0.0.1
  * @apiName faceRecognition
  * @apiGroup Face Recognition
  * @apiDescription 人臉辦識，可以選擇可信度，及最大回傳數目，結果會回傳0-max個辨識結果高於可信度值的用戶資訊
  * @apiParam {String} image 臉部圖片(base64-encoded)
  * @apiParam {Number}  threshold  可信度(0.0-1.0),數字愈高可信度愈高
  * @apiParam {Numnber} max  最多回傳用戶數目 (目前只支援1個)
  *
  * @apiSuccess {Number} code  錯誤代碼
  *                 0:SUCCESS(成功)
  *                 1:INVALID_PARAMETERS(參數缺少或錯誤
  *                 3:USER_NOT_EXIST(用戶不存在)
  *                 4:INVALID_IMAGE(圖片格式不符)
  * @apiSuccess {Object[]} userList 回傳用戶資訊陣列(物件參考/api/user/info)
  */
  function faceRecognition() { return; }
