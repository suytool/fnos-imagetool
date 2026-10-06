/* 图片工具 - 前端逻辑 */
'use strict';

/* ==================== 模板定义（与后端一致） ==================== */
function uniform(rows, cols) {
  const cells = [];
  for (let r = 0; r < rows; r++) for (let c = 0; c < cols; c++)
    cells.push({ row: r, col: c, rs: 1, cs: 1, img: true });
  return cells;
}

const COLLAGE = {
  classic9: [{ id: 'classic9-a', name: '经典九宫格 3×3', rows: 3, cols: 3, cells: uniform(3, 3) }],
  '2': [
    { id: '2-a', name: '左右各一', rows: 1, cols: 2, cells: uniform(1, 2) },
    { id: '2-b', name: '上下各一', rows: 2, cols: 1, cells: uniform(2, 1) },
    { id: '2-c', name: '左大右小', rows: 2, cols: 2, cells: [
      { row: 0, col: 0, rs: 2, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: false }] },
    { id: '2-d', name: '上大下小', rows: 2, cols: 2, cells: [
      { row: 0, col: 0, rs: 1, cs: 2, img: true }, { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: false }] },
  ],
  '3': [
    { id: '3-a', name: '横排三图', rows: 1, cols: 3, cells: uniform(1, 3) },
    { id: '3-b', name: '竖排三图', rows: 3, cols: 1, cells: uniform(3, 1) },
    { id: '3-c', name: '上1下2', rows: 2, cols: 2, cells: [
      { row: 0, col: 0, rs: 1, cs: 2, img: true }, { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }] },
    { id: '3-d', name: '上2下1', rows: 2, cols: 2, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 0, rs: 1, cs: 2, img: true }] },
    { id: '3-e', name: '左1右2', rows: 2, cols: 2, cells: [
      { row: 0, col: 0, rs: 2, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }] },
    { id: '3-f', name: '左2右1', rows: 2, cols: 2, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 2, cs: 1, img: true }] },
  ],
  '4': [
    { id: '4-a', name: '田字格 2×2', rows: 2, cols: 2, cells: uniform(2, 2) },
    { id: '4-b', name: '横排四图', rows: 1, cols: 4, cells: uniform(1, 4) },
    { id: '4-c', name: '竖排四图', rows: 4, cols: 1, cells: uniform(4, 1) },
    { id: '4-d', name: '上1下3', rows: 2, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 3, img: true }, { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 1, img: true }] },
    { id: '4-e', name: '左1右3', rows: 3, cols: 2, cells: [
      { row: 0, col: 0, rs: 3, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }] },
  ],
  '5': [
    { id: '5-a', name: '上2下3', rows: 2, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 1, img: true }] },
    { id: '5-b', name: '上3下2', rows: 2, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 0, col: 2, rs: 1, cs: 1, img: true }, { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }] },
    { id: '5-c', name: '上1下4', rows: 2, cols: 4, cells: [
      { row: 0, col: 0, rs: 1, cs: 4, img: true }, { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 1, img: true }, { row: 1, col: 3, rs: 1, cs: 1, img: true }] },
    { id: '5-d', name: '2-1-2 三段', rows: 3, cols: 2, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 0, rs: 1, cs: 2, img: true }, { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }] },
  ],
  '6': [
    { id: '6-a', name: '2行3列', rows: 2, cols: 3, cells: uniform(2, 3) },
    { id: '6-b', name: '3行2列', rows: 3, cols: 2, cells: uniform(3, 2) },
    { id: '6-c', name: '2-2-2', rows: 3, cols: 2, cells: uniform(3, 2) },
    { id: '6-d', name: '上3中1下2', rows: 3, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 0, col: 2, rs: 1, cs: 1, img: true },
      { row: 1, col: 0, rs: 1, cs: 3, img: true },
      { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }, { row: 2, col: 2, rs: 1, cs: 1, img: false }] },
  ],
  '7': [
    { id: '7-a', name: '2-3-2', rows: 3, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true },
      { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 1, img: true },
      { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }] },
    { id: '7-b', name: '上1中3下3', rows: 3, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 3, img: true },
      { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 1, img: true },
      { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }, { row: 2, col: 2, rs: 1, cs: 1, img: true }] },
    { id: '7-c', name: '3-1-3', rows: 3, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 0, col: 2, rs: 1, cs: 1, img: true },
      { row: 1, col: 0, rs: 1, cs: 3, img: true },
      { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }, { row: 2, col: 2, rs: 1, cs: 1, img: true }] },
  ],
  '8': [
    { id: '8-a', name: '2行4列', rows: 2, cols: 4, cells: uniform(2, 4) },
    { id: '8-b', name: '4行2列', rows: 4, cols: 2, cells: uniform(4, 2) },
    { id: '8-c', name: '3-2-3', rows: 3, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true }, { row: 0, col: 2, rs: 1, cs: 1, img: true },
      { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true },
      { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }, { row: 2, col: 2, rs: 1, cs: 1, img: true }] },
    { id: '8-d', name: '2-4-2', rows: 3, cols: 4, cells: [
      { row: 0, col: 0, rs: 1, cs: 1, img: true }, { row: 0, col: 1, rs: 1, cs: 1, img: true },
      { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 1, img: true }, { row: 1, col: 3, rs: 1, cs: 1, img: true },
      { row: 2, col: 2, rs: 1, cs: 1, img: true }, { row: 2, col: 3, rs: 1, cs: 1, img: true }] },
  ],
  '9': [
    { id: '9-a', name: '三行三列', rows: 3, cols: 3, cells: uniform(3, 3) },
    { id: '9-b', name: '大图置顶', rows: 3, cols: 3, cells: [
      { row: 0, col: 0, rs: 1, cs: 2, img: true }, { row: 0, col: 2, rs: 1, cs: 1, img: true },
      { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 1, img: true },
      { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }, { row: 2, col: 2, rs: 1, cs: 1, img: true }] },
    { id: '9-c', name: '2-3-4', rows: 3, cols: 4, cells: [
      { row: 0, col: 0, rs: 1, cs: 2, img: true }, { row: 0, col: 2, rs: 1, cs: 2, img: true },
      { row: 1, col: 0, rs: 1, cs: 1, img: true }, { row: 1, col: 1, rs: 1, cs: 1, img: true }, { row: 1, col: 2, rs: 1, cs: 2, img: true },
      { row: 2, col: 0, rs: 1, cs: 1, img: true }, { row: 2, col: 1, rs: 1, cs: 1, img: true }, { row: 2, col: 2, rs: 1, cs: 1, img: true }, { row: 2, col: 3, rs: 1, cs: 1, img: true }] },
  ],
};

const SPLIT = {
  'split-9': { id: 'split-9', name: '1拆9 (3×3)', rows: 3, cols: 3 },
  'split-6': { id: 'split-6', name: '1拆6 (2×3)', rows: 2, cols: 3 },
  'split-4': { id: 'split-4', name: '1拆4 (2×2)', rows: 2, cols: 2 },
};

/* 菜单结构 */
const MENU = [
  {
    group: '拼图模式',
    items: [
      { label: '经典九宫格', key: 'classic9' },
      { label: '2图模式', key: '2' },
      { label: '3图模式', key: '3' },
      { label: '4图模式', key: '4' },
      { label: '5图模式', key: '5' },
      { label: '6图模式', key: '6' },
      { label: '7图模式', key: '7' },
      { label: '8图模式', key: '8' },
      { label: '9图模式', key: '9' },
      { label: '自由拼图', key: 'free' },
    ],
  },
  {
    group: '分图模式',
    items: [
      { label: '1拆9 (3x3)', key: 'split-9' },
      { label: '1拆6 (2x3)', key: 'split-6' },
      { label: '1拆4 (2x2)', key: 'split-4' },
      { label: '自由分图', key: 'freesplit' },
    ],
  },
];

/* ==================== 状态 ==================== */
const state = {
  mode: null,        // classic9/2../9/free/split-9/split-6/split-4/freesplit
  variant: null,     // 拼图布局变体 id
  libImages: [],     // {file, url, w, h}
  cells: [],         // 当前模板格子图片索引: [libIdx|null]
  cellViews: [],     // 每格视图调整: {zoom, dx, dy}
  selectedCell: null,// 当前选中的有图格子序号
  layout: null,      // 当前拼图布局
  splitRows: 3, splitCols: 3,
  splitFile: null,
  lastCollageBlob: null,
};

/* ==================== 工具函数 ==================== */
const $ = (id) => document.getElementById(id);

function isSplitMode(mode) {
  return mode && (mode.startsWith('split') || mode === 'freesplit');
}

function showError(id, msg) { $(id).textContent = msg; }

function downloadBlob(blob, filename) {
  const a = document.createElement('a');
  a.href = URL.createObjectURL(blob);
  a.download = filename;
  document.body.appendChild(a);
  a.click();
  setTimeout(() => { URL.revokeObjectURL(a.href); a.remove(); }, 2000);
}

/* ==================== 标签页切换 ==================== */
document.querySelectorAll('.nav-tabs button').forEach((btn) => {
  btn.addEventListener('click', () => {
    document.querySelectorAll('.nav-tabs button').forEach((b) => b.classList.remove('active'));
    btn.classList.add('active');
    const tab = btn.dataset.tab;
    document.querySelectorAll('.page').forEach((p) => p.classList.remove('active'));
    $('page-' + tab).classList.add('active');
    if (tab === 'collage') {
      document.querySelector('.split-page').classList.add('active');
    }
  });
});

/* ==================== 通用上传区（拖拽/点击） ==================== */
function bindUploadZone(zone, input, onFile) {
  zone.addEventListener('click', () => input.click());
  input.addEventListener('change', () => {
    if (input.files.length) { onFile(input.files); input.value = ''; }
  });
  ['dragover', 'dragenter'].forEach((ev) => zone.addEventListener(ev, (e) => {
    e.preventDefault(); zone.classList.add('dragover');
  }));
  ['dragleave', 'drop'].forEach((ev) => zone.addEventListener(ev, (e) => {
    e.preventDefault(); zone.classList.remove('dragover');
  }));
  zone.addEventListener('drop', (e) => {
    if (e.dataTransfer && e.dataTransfer.files.length) onFile(e.dataTransfer.files);
  });
}

/* ==================== 模板下拉菜单 ==================== */
const dropdown = $('tpl-dropdown');
$('tpl-btn').addEventListener('click', (e) => {
  e.stopPropagation();
  dropdown.classList.toggle('open');
});
document.addEventListener('click', () => dropdown.classList.remove('open'));

function buildMenu() {
  const menu = $('tpl-menu');
  menu.innerHTML = '';
  MENU.forEach((g) => {
    const gDiv = document.createElement('div');
    gDiv.className = 'tpl-group';
    gDiv.innerHTML = `<div class="tpl-group-title">${g.group}</div>`;
    g.items.forEach((it) => {
      const isSplit = it.key.startsWith('split') || it.key === 'freesplit';
      const hasVariants = !isSplit && COLLAGE[it.key] && COLLAGE[it.key].length > 1;
      const item = document.createElement('div');
      item.className = 'tpl-item';
      item.dataset.key = it.key;
      item.innerHTML = `<span>${it.label}</span>${hasVariants ? '<span class="caret">▸</span>' : ''}`;
      item.addEventListener('click', (e) => {
        e.stopPropagation();
        handleMenuPick(it.key, item);
      });
      gDiv.appendChild(item);

      // 变体区
      if (hasVariants) {
        const vDiv = document.createElement('div');
        vDiv.className = 'variants';
        vDiv.id = 'variants-' + it.key;
        const vg = document.createElement('div');
        vg.className = 'v-grid';
        COLLAGE[it.key].forEach((v) => {
          const b = document.createElement('button');
          b.className = 'v-btn';
          b.textContent = v.name;
          b.dataset.vid = v.id;
          b.addEventListener('click', (e2) => {
            e2.stopPropagation();
            selectCollage(it.key, v.id);
            dropdown.classList.remove('open');
          });
          vg.appendChild(b);
        });
        vDiv.appendChild(vg);
        gDiv.appendChild(vDiv);
      }
    });
    menu.appendChild(gDiv);
  });
}

function handleMenuPick(key, item) {
  // 高亮
  document.querySelectorAll('.tpl-item').forEach((i) => i.classList.remove('selected'));
  item.classList.add('selected');
  document.querySelectorAll('.variants').forEach((v) => v.classList.remove('open'));

  const isSplit = key.startsWith('split') || key === 'freesplit';
  if (!isSplit && COLLAGE[key] && COLLAGE[key].length > 1) {
    // 展开变体，选中第一个
    const vDiv = $('variants-' + key);
    vDiv.classList.add('open');
    selectCollage(key, COLLAGE[key][0].id);
  } else if (!isSplit) {
    selectCollage(key, COLLAGE[key] ? COLLAGE[key][0].id : null);
  } else {
    selectSplit(key);
  }
}

/* 选择拼图模板 */
function selectCollage(key, variant) {
  const layout = COLLAGE[key].find((v) => v.id === variant) || COLLAGE[key][0];
  state.mode = key;
  state.variant = layout.id;
  state.layout = layout;
  state.splitFile = null;
  $('free-collage-params').classList.toggle('open', key === 'free');
  $('free-split-params').classList.remove('open');
  $('tpl-selected').textContent = key === 'free'
    ? `自由拼图 ${layout.rows}×${layout.cols}`
    : `${menuLabel(key)} · ${layout.name}`;
  document.querySelectorAll('.v-btn').forEach((b) => b.classList.toggle('active', b.dataset.vid === layout.id));
  renderCollageCanvas();
  dropdown.classList.remove('open');
}

/* 选择分图模板 */
function selectSplit(key) {
  state.mode = key;
  state.layout = null;
  state.cells = [];
  if (key === 'freesplit') {
    $('free-split-params').classList.add('open');
    $('free-collage-params').classList.remove('open');
    state.splitRows = clampInt(+$('free-split-rows').value, 1, 10);
    state.splitCols = clampInt(+$('free-split-cols').value, 1, 10);
    $('tpl-selected').textContent = `自由分图 ${state.splitRows}×${state.splitCols}`;
  } else {
    const st = SPLIT[key];
    state.splitRows = st.rows;
    state.splitCols = st.cols;
    $('free-split-params').classList.remove('open');
    $('free-collage-params').classList.remove('open');
    $('tpl-selected').textContent = st.name;
  }
  renderSplitCanvas();
  dropdown.classList.remove('open');
}

function menuLabel(key) {
  const flat = [];
  MENU.forEach((g) => g.items.forEach((i) => flat.push(i)));
  const it = flat.find((i) => i.key === key);
  return it ? it.label : key;
}

function clampInt(v, lo, hi) { return Math.max(lo, Math.min(hi, v | 0)); }

/* 自由拼图行列变化 */
$('free-rows').addEventListener('change', () => {
  if (state.mode === 'free') {
    state.layout.rows = clampInt(+$('free-rows').value, 1, 7);
    state.layout.cells = uniform(state.layout.rows, state.layout.cols);
    state.cells = state.cells.slice(0, state.layout.cells.filter(c => c.img).length);
    $('tpl-selected').textContent = `自由拼图 ${state.layout.rows}×${state.layout.cols}`;
    renderCollageCanvas();
  }
});
$('free-cols').addEventListener('change', () => {
  if (state.mode === 'free') {
    state.layout.cols = clampInt(+$('free-cols').value, 1, 7);
    state.layout.cells = uniform(state.layout.rows, state.layout.cols);
    state.cells = state.cells.slice(0, state.layout.cells.filter(c => c.img).length);
    $('tpl-selected').textContent = `自由拼图 ${state.layout.rows}×${state.layout.cols}`;
    renderCollageCanvas();
  }
});
$('free-split-rows').addEventListener('change', () => {
  if (state.mode === 'freesplit') {
    state.splitRows = clampInt(+$('free-split-rows').value, 1, 10);
    $('tpl-selected').textContent = `自由分图 ${state.splitRows}×${state.splitCols}`;
    renderSplitOverlay();
  }
});
$('free-split-cols').addEventListener('change', () => {
  if (state.mode === 'freesplit') {
    state.splitCols = clampInt(+$('free-split-cols').value, 1, 10);
    $('tpl-selected').textContent = `自由分图 ${state.splitRows}×${state.splitCols}`;
    renderSplitOverlay();
  }
});

/* ==================== 素材库 ==================== */
function addLibImage(file) {
  return new Promise((resolve) => {
    const url = URL.createObjectURL(file);
    const im = new Image();
    im.onload = () => {
      state.libImages.push({ file, url, w: im.naturalWidth, h: im.naturalHeight });
      resolve(state.libImages.length - 1);
    };
    im.onerror = () => {
      state.libImages.push({ file, url, w: 0, h: 0 });
      resolve(state.libImages.length - 1);
    };
    im.src = url;
  });
}

function addFilesToLib(files, cb) {
  const list = Array.from(files).filter((f) => f.type.startsWith('image/'));
  let firstIdx = null;
  Promise.all(list.map((f) => addLibImage(f).then((i) => { if (firstIdx === null) firstIdx = i; })))
    .then(() => {
      renderLibrary();
      if (cb && firstIdx !== null) cb(firstIdx);
    });
}

bindUploadZone($('library-zone'), $('lib-file'), (files) => {
  addFilesToLib(files, null);
});

function renderLibrary() {
  const box = $('lib-items');
  box.innerHTML = '';
  if (!state.libImages.length) {
    box.innerHTML = '<div class="lib-empty">暂无素材，请上传图片</div>';
    return;
  }
  state.libImages.forEach((it, idx) => {
    const d = document.createElement('div');
    d.className = 'lib-item';
    d.draggable = true;
    d.dataset.idx = idx;
    const img = document.createElement('img');
    img.src = it.url;
    d.appendChild(img);
    const rm = document.createElement('button');
    rm.className = 'remove';
    rm.textContent = '×';
    rm.addEventListener('click', (e) => {
      e.stopPropagation();
      URL.revokeObjectURL(it.url);
      state.libImages.splice(idx, 1);
      // 清除画布中对该素材的引用
      state.cells = state.cells.map((c) => (c === idx ? null : (c !== null && c > idx ? c - 1 : c)));
      renderLibrary();
      renderCollageCanvas();
    });
    d.appendChild(rm);
    d.addEventListener('dragstart', (e) => {
      e.dataTransfer.setData('application/x-lib-idx', String(idx));
      d.classList.add('dragging');
    });
    d.addEventListener('dragend', () => d.classList.remove('dragging'));
    box.appendChild(d);
  });
}

/* ==================== 拼图画布 ==================== */
function renderCollageCanvas() {
  const layout = state.layout;
  const body = $('canvas-body');
  const hint = $('canvas-hint');
  const gridWrap = $('grid-wrap');
  const splitDrop = $('split-drop');
  const splitPreview = $('split-preview');
  const runSplit = $('run-split');
  const runCollage = $('run-collage');

  hint.style.display = layout ? 'none' : 'block';
  gridWrap.style.display = layout ? 'inline-block' : 'none';
  splitDrop.style.display = 'none';
  splitPreview.style.display = 'none';
  runSplit.style.display = 'none';
  runCollage.style.display = 'inline-block';
  $('canvas-mode-badge').textContent = layout ? (layout.rows + '×' + layout.cols + ' · ' + layout.cells.filter(c => c.img).length + '图') : '未选择模板';

  if (!layout) {
    gridWrap.innerHTML = '';
    state.cells = [];
    return;
  }

  // 初始化 cells / cellViews
  const n = layout.cells.filter((c) => c.img).length;
  state.cells = state.cells.slice(0, n);
  while (state.cells.length < n) state.cells.push(null);
  state.cellViews = state.cellViews.slice(0, n);
  while (state.cellViews.length < n) state.cellViews.push({ zoom: 1, dx: 0, dy: 0 });
  if (state.selectedCell !== null && state.selectedCell >= n) state.selectedCell = null;

  const gap = +$('cfg-gap').value || 10;
  const grid = document.createElement('div');
  grid.className = 'grid';
  grid.style.gap = gap + 'px';
  grid.style.gridTemplateColumns = `repeat(${layout.cols}, 1fr)`;
  grid.style.gridTemplateRows = `repeat(${layout.rows}, 1fr)`;
  grid.style.width = '560px';
  grid.style.aspectRatio = `${layout.cols} / ${layout.rows}`;

  let imgIdx = 0;
  layout.cells.forEach((cell, ci) => {
    const el = document.createElement('div');
    el.className = 'cell' + (cell.img ? '' : ' empty');
    if (cell.img) {
      el.style.gridRow = `${cell.row + 1} / span ${cell.rs}`;
      el.style.gridColumn = `${cell.col + 1} / span ${cell.cs}`;
      el.dataset.cell = imgIdx;
      const libIdx = state.cells[imgIdx];
      if (libIdx !== null && libIdx !== undefined && state.libImages[libIdx]) {
        el.classList.add('has-img');
        const it = state.libImages[libIdx];
        const im = document.createElement('img');
        im.className = 'cell-img';
        im.src = it.url;
        im.draggable = false;
        im.addEventListener('load', () => { if (el.isConnected) layoutCellImage(el, im, state.cellViews[imgIdx], it.w, it.h); });
        el.appendChild(im);
        const rm = document.createElement('button');
        rm.className = 'remove-img';
        rm.textContent = '×';
        rm.addEventListener('click', (e) => {
          e.stopPropagation();
          state.cells[+el.dataset.cell] = null;
          state.cellViews[+el.dataset.cell] = { zoom: 1, dx: 0, dy: 0 };
          if (state.selectedCell === +el.dataset.cell) selectCell(null);
          renderCollageCanvas();
        });
        el.appendChild(rm);
        // 视图定位 + 交互
        layoutCellImage(el, im, state.cellViews[imgIdx], it.w, it.h);
        bindCellInteractions(el, imgIdx);
      } else {
        const no = document.createElement('span');
        no.className = 'cell-no';
        no.textContent = imgIdx + 1;
        el.appendChild(no);
        el.addEventListener('click', () => {
          if (state.libImages.length) {
            state.cells[+el.dataset.cell] = 0;
            renderCollageCanvas();
          }
        });
      }
      imgIdx++;
      el.addEventListener('dragover', (e) => { e.preventDefault(); el.classList.add('dragover'); });
      el.addEventListener('dragleave', () => el.classList.remove('dragover'));
      el.addEventListener('drop', (e) => {
        e.preventDefault();
        el.classList.remove('dragover');
        const idx = e.dataTransfer.getData('application/x-lib-idx');
        if (idx !== '') {
          const i = +idx;
          if (state.libImages[i]) {
            state.cells[+el.dataset.cell] = i;
            state.cellViews[+el.dataset.cell] = { zoom: 1, dx: 0, dy: 0 };
          }
        } else if (e.dataTransfer.files.length) {
          addFilesToLib(e.dataTransfer.files, (newIdx) => {
            state.cells[+el.dataset.cell] = newIdx;
            state.cellViews[+el.dataset.cell] = { zoom: 1, dx: 0, dy: 0 };
            renderCollageCanvas();
          });
        }
        renderCollageCanvas();
      });
    } else {
      el.style.gridRow = `${cell.row + 1} / span ${cell.rs}`;
      el.style.gridColumn = `${cell.col + 1} / span ${cell.cs}`;
      el.style.border = 'none';
      el.style.background = 'transparent';
    }
    grid.appendChild(el);
  });

  gridWrap.innerHTML = '';
  gridWrap.appendChild(grid);
  $('auto-fill').disabled = false;
  requestAnimationFrame(() => { layoutAllCellImages(); syncCellEditUI(); });
}

/* ---------- 格子视图（缩放/平移） ---------- */

// 夹紧视图参数（与后端 fillCoverView 的夹紧算法一致）
function clampView(v, cw, ch, sw, sh) {
  if (!sw || !sh || !cw || !ch) { v.zoom = 1; v.dx = 0; v.dy = 0; v._maxDx = 100; v._maxDy = 100; return v; }
  const z = Math.max(0.5, Math.min(5, v.zoom));
  const s0 = Math.max(cw / sw, ch / sh);
  const worldW = sw * s0, worldH = sh * s0;
  let ww = cw / z, wh = ch / z, maxDx, maxDy;
  if (z < 1) {
    ww = cw; wh = ch;
    maxDx = 50 * (worldW / z / ww - 1);
    maxDy = 50 * (worldH / z / wh - 1);
  } else {
    maxDx = 50 * (worldW / ww - 1);
    maxDy = 50 * (worldH / wh - 1);
  }
  if (maxDx < 0) maxDx = 0;
  if (maxDy < 0) maxDy = 0;
  maxDx = Math.min(maxDx, 100);
  maxDy = Math.min(maxDy, 100);
  v.zoom = z;
  v.dx = Math.max(-maxDx, Math.min(maxDx, v.dx));
  v.dy = Math.max(-maxDy, Math.min(maxDy, v.dy));
  v._maxDx = maxDx;
  v._maxDy = maxDy;
  return v;
}

// 按视图像素定位格子内图片
function layoutCellImage(cellEl, imgEl, view, sw, sh) {
  const rect = cellEl.getBoundingClientRect();
  const cw = rect.width, ch = rect.height;
  if (!cw || !ch || !sw || !sh) return;
  const z = view.zoom;
  const s0 = Math.max(cw / sw, ch / sh);
  let imgW, imgH, winW, winH;
  if (z >= 1) { imgW = sw * s0 * z; imgH = sh * s0 * z; winW = cw / z; winH = ch / z; }
  else { imgW = sw * s0 / z; imgH = sh * s0 / z; winW = cw; winH = ch; }
  let left = (cw - imgW) / 2 - (view.dx / 100) * winW;
  let top = (ch - imgH) / 2 - (view.dy / 100) * winH;
  left = Math.max(cw - imgW, Math.min(0, left));
  top = Math.max(ch - imgH, Math.min(0, top));
  imgEl.style.width = Math.round(imgW) + 'px';
  imgEl.style.height = Math.round(imgH) + 'px';
  imgEl.style.left = Math.round(left) + 'px';
  imgEl.style.top = Math.round(top) + 'px';
}

function layoutAllCellImages() {
  const grid = $('grid-wrap');
  if (!grid || !grid.firstChild) return;
  grid.querySelectorAll('.cell.has-img').forEach((el) => {
    const ci = +el.dataset.cell;
    const libIdx = state.cells[ci];
    const im = el.querySelector('img.cell-img');
    if (im && libIdx !== null && libIdx !== undefined && state.libImages[libIdx]) {
      const it = state.libImages[libIdx];
      layoutCellImage(el, im, state.cellViews[ci], it.w, it.h);
    }
  });
}

// 格子交互：点击选中、拖动平移、滚轮缩放
function bindCellInteractions(cellEl, cellIdx) {
  cellEl.addEventListener('click', (e) => {
    if (e.target.classList.contains('remove-img')) return;
    if (state.cells[cellIdx] !== null && state.cells[cellIdx] !== undefined) {
      selectCell(cellIdx);
    }
  });
  // 拖动平移
  let drag = null;
  cellEl.addEventListener('mousedown', (e) => {
    if (e.button !== 0) return;
    if (state.cells[cellIdx] === null || state.cells[cellIdx] === undefined) return;
    if (e.target.classList.contains('remove-img')) return;
    e.preventDefault();
    const it = state.libImages[state.cells[cellIdx]];
    const rect = cellEl.getBoundingClientRect();
    const cw = rect.width, ch = rect.height;
    const view = state.cellViews[cellIdx];
    const z = view.zoom;
    const winW = z >= 1 ? cw / z : cw;
    const winH = z >= 1 ? ch / z : ch;
    drag = { sx: e.clientX, sy: e.clientY, dx0: view.dx, dy0: view.dy, winW, winH, it, cw, ch };
    cellEl.classList.add('dragging');
  });
  window.addEventListener('mousemove', (e) => {
    if (!drag) return;
    const view = state.cellViews[cellIdx];
    view.dx = drag.dx0 - ((e.clientX - drag.sx) / drag.winW) * 100;
    view.dy = drag.dy0 - ((e.clientY - drag.sy) / drag.winH) * 100;
    clampView(view, drag.cw, drag.ch, drag.it.w, drag.it.h);
    const im = cellEl.querySelector('img.cell-img');
    if (im) layoutCellImage(cellEl, im, view, drag.it.w, drag.it.h);
    syncCellEditUI();
  });
  window.addEventListener('mouseup', () => {
    if (!drag) return;
    cellEl.classList.remove('dragging');
    drag = null;
  });
  // 滚轮缩放
  cellEl.addEventListener('wheel', (e) => {
    if (state.cells[cellIdx] === null || state.cells[cellIdx] === undefined) return;
    e.preventDefault();
    e.stopPropagation();
    const it = state.libImages[state.cells[cellIdx]];
    const view = state.cellViews[cellIdx];
    const rect = cellEl.getBoundingClientRect();
    view.zoom *= e.deltaY < 0 ? 1.1 : 1 / 1.1;
    clampView(view, rect.width, rect.height, it.w, it.h);
    const im = cellEl.querySelector('img.cell-img');
    if (im) layoutCellImage(cellEl, im, view, it.w, it.h);
    syncCellEditUI();
  }, { passive: false });
}

/* 选中格子：显示调整工具栏 */
function selectCell(cellIdx) {
  state.selectedCell = cellIdx;
  document.querySelectorAll('.cell').forEach((c) => c.classList.remove('selected'));
  const grid = $('grid-wrap');
  if (grid) {
    grid.querySelectorAll('.cell.has-img').forEach((c) => {
      if (+c.dataset.cell === cellIdx) c.classList.add('selected');
    });
  }
  syncCellEditUI();
}

function syncCellEditUI() {
  const bar = $('cell-edit');
  const ci = state.selectedCell;
  if (ci === null || ci === undefined || !state.layout || state.cells[ci] === null || state.cells[ci] === undefined || !state.libImages[state.cells[ci]]) {
    bar.style.display = 'none';
    state.selectedCell = null;
    return;
  }
  bar.style.display = 'flex';
  $('ce-idx').textContent = ci + 1;
  const v = state.cellViews[ci];
  $('ce-zoom').value = Math.round(v.zoom * 100);
  $('ce-zoom-val').textContent = Math.round(v.zoom * 100) + '%';
  const maxDx = Math.round(v._maxDx !== undefined ? v._maxDx : 100);
  const maxDy = Math.round(v._maxDy !== undefined ? v._maxDy : 100);
  $('ce-dx').min = -maxDx;
  $('ce-dx').max = maxDx;
  $('ce-dy').min = -maxDy;
  $('ce-dy').max = maxDy;
  $('ce-dx').value = Math.round(v.dx);
  $('ce-dx-val').textContent = Math.round(v.dx);
  $('ce-dy').value = Math.round(v.dy);
  $('ce-dy-val').textContent = Math.round(v.dy);
}

/* 工具栏滑块事件 */
$('ce-zoom').addEventListener('input', () => {
  const ci = state.selectedCell;
  if (ci === null) return;
  state.cellViews[ci].zoom = (+$('ce-zoom').value) / 100;
  const it = state.libImages[state.cells[ci]];
  const cellEl = document.querySelector(`.cell[data-cell="${ci}"]`);
  if (cellEl && it) {
    clampView(state.cellViews[ci], cellEl.getBoundingClientRect().width, cellEl.getBoundingClientRect().height, it.w, it.h);
    const im = cellEl.querySelector('img.cell-img');
    if (im) layoutCellImage(cellEl, im, state.cellViews[ci], it.w, it.h);
  }
  syncCellEditUI();
});
$('ce-dx').addEventListener('input', () => {
  const ci = state.selectedCell;
  if (ci === null) return;
  state.cellViews[ci].dx = +$('ce-dx').value;
  applySelectedView();
});
$('ce-dy').addEventListener('input', () => {
  const ci = state.selectedCell;
  if (ci === null) return;
  state.cellViews[ci].dy = +$('ce-dy').value;
  applySelectedView();
});
function applySelectedView() {
  const ci = state.selectedCell;
  if (ci === null) return;
  const it = state.libImages[state.cells[ci]];
  const cellEl = document.querySelector(`.cell[data-cell="${ci}"]`);
  if (cellEl && it) {
    clampView(state.cellViews[ci], cellEl.getBoundingClientRect().width, cellEl.getBoundingClientRect().height, it.w, it.h);
    const im = cellEl.querySelector('img.cell-img');
    if (im) layoutCellImage(cellEl, im, state.cellViews[ci], it.w, it.h);
  }
  syncCellEditUI();
}
$('ce-reset').addEventListener('click', () => {
  const ci = state.selectedCell;
  if (ci === null) return;
  state.cellViews[ci] = { zoom: 1, dx: 0, dy: 0 };
  applySelectedView();
});

/* 自动填充 */
$('auto-fill').addEventListener('click', () => {
  let libIdx = 0;
  state.cells = state.cells.map((c) => {
    if (c === null && libIdx < state.libImages.length) return libIdx++;
    return c;
  });
  renderCollageCanvas();
});

/* 清空画布 */
$('clear-canvas').addEventListener('click', () => {
  state.cells = [];
  state.cellViews = [];
  selectCell(null);
  $('cell-edit').style.display = 'none';
  state.splitFile = null;
  $('split-preview').style.display = 'none';
  $('split-preview-img').src = '';
  renderCollageCanvas();
  renderSplitCanvas();
  hideResult();
});

/* ==================== 分图画布 ==================== */
function renderSplitCanvas() {
  const body = $('canvas-body');
  const hint = $('canvas-hint');
  const gridWrap = $('grid-wrap');
  const splitDrop = $('split-drop');
  const splitPreview = $('split-preview');
  const runSplit = $('run-split');
  const runCollage = $('run-collage');

  hint.style.display = 'none';
  gridWrap.style.display = 'none';
  splitDrop.style.display = state.splitFile ? 'none' : 'block';
  splitPreview.style.display = state.splitFile ? 'block' : 'none';
  runSplit.style.display = 'inline-block';
  runCollage.style.display = 'none';
  $('canvas-mode-badge').textContent = `${state.splitRows}×${state.splitCols} 拆分`;
  $('auto-fill').disabled = true;
  renderSplitOverlay();
}

function renderSplitOverlay() {
  const ov = $('split-overlay');
  ov.style.setProperty('--cw', (100 / state.splitCols) + '%');
  ov.style.setProperty('--ch', (100 / state.splitRows) + '%');
  ov.style.backgroundSize = `${100 / state.splitCols}% ${100 / state.splitRows}%`;
  if (state.splitFile) {
    const img = $('split-preview-img');
    // 等待图片加载后设置 overlay 尺寸
    img.onload = () => {
      ov.style.width = img.clientWidth + 'px';
      ov.style.height = img.clientHeight + 'px';
    };
    ov.style.width = img.clientWidth + 'px';
    ov.style.height = img.clientHeight + 'px';
  }
}

bindUploadZone($('split-drop'), $('split-file'), (files) => {
  const f = Array.from(files).find((x) => x.type.startsWith('image/'));
  if (!f) return;
  state.splitFile = f;
  $('split-preview-img').src = URL.createObjectURL(f);
  renderSplitCanvas();
});

/* ==================== 生成拼图 ==================== */
$('run-collage').addEventListener('click', async () => {
  const layout = state.layout;
  if (!layout) { showError('collage-error', '请先选择拼图模板'); return; }
  const need = layout.cells.filter((c) => c.img).length;
  const filled = state.cells.filter((c) => c !== null && c !== undefined).length;
  if (filled < need) { showError('collage-error', `模板需要 ${need} 张图片，当前已放入 ${filled} 张`); return; }

  const fd = new FormData();
  const views = [];
  layout.cells.forEach((cell, i) => {
    if (cell.img && state.cells[i] !== null) {
      fd.append('images', state.libImages[state.cells[i]].file);
      const v = state.cellViews[i] || { zoom: 1, dx: 0, dy: 0 };
      views.push({ zoom: Math.round(v.zoom * 100) / 100, dx: Math.round(v.dx * 10) / 10, dy: Math.round(v.dy * 10) / 10 });
    }
  });
  fd.append('cells', JSON.stringify(views));
  if (state.mode === 'free') {
    fd.append('mode', 'free');
    fd.append('rows', layout.rows);
    fd.append('cols', layout.cols);
  } else {
    fd.append('mode', state.mode);
    fd.append('variant', layout.id);
  }
  fd.append('gap', $('cfg-gap').value || '10');
  fd.append('bg', $('cfg-bg').value || '#ffffff');
  fd.append('format', $('cfg-format').value);
  fd.append('quality', $('cfg-quality').value || '90');

  showResultLoading();

  try {
    const blob = await postForm('/api/collage', fd);
    if (!blob) return;
    showResultImage(blob, '拼图结果', false);
    state.lastCollageBlob = blob;
  } catch (e) {
    showError('collage-error', e.message);
  } finally {
    $('collage-loading').style.display = 'none';
  }
});

$('collage-download').addEventListener('click', () => {
  if (!state.lastCollageBlob) return;
  if (state.mode && isSplitMode(state.mode)) {
    downloadBlob(state.lastCollageBlob, `split_${state.splitRows}x${state.splitCols}.zip`);
  } else {
    downloadBlob(state.lastCollageBlob, 'collage.' + ($('cfg-format').value === 'png' ? 'png' : 'jpg'));
  }
});

/* ==================== 开始分图 ==================== */
$('run-split').addEventListener('click', async () => {
  if (!state.splitFile) { showError('collage-error', '请先上传要拆分的图片'); return; }
  const fd = new FormData();
  fd.append('image', state.splitFile);
  fd.append('rows', state.splitRows);
  fd.append('cols', state.splitCols);
  fd.append('format', $('cfg-format').value);
  fd.append('quality', $('cfg-quality').value || '90');

  showResultLoading();

  try {
    const blob = await postForm('/api/split', fd);
    if (!blob) return;
    showResultImage(blob, '分图结果', true);
    state.lastCollageBlob = blob;
  } catch (e) {
    showError('collage-error', e.message);
  } finally {
    $('collage-loading').style.display = 'none';
  }
});

/* 结果展示 */
function showResultLoading() {
  const box = $('collage-result');
  box.classList.add('show');
  $('collage-loading').style.display = 'block';
  $('collage-result-body').style.display = 'none';
  $('collage-close').style.display = 'none';
  showError('collage-error', '');
}

function showResultImage(blob, title, isZip) {
  const img = $('collage-result-img');
  const meta = $('collage-result-meta');
  if (isZip) {
    img.style.display = 'none';
    meta.innerHTML = `<b>${title}</b> · 打包 zip · ${formatSize(blob.size)}<br>共 ${state.splitRows}×${state.splitCols} = ${state.splitRows * state.splitCols} 片，已按顺序编号`;
  } else {
    img.style.display = 'block';
    const url = URL.createObjectURL(blob);
    img.onload = () => {
      meta.innerHTML = `<b>${title}</b> · ${img.naturalWidth} × ${img.naturalHeight} px · ${formatSize(blob.size)}`;
    };
    meta.innerHTML = `<b>${title}</b> · 处理中...`;
    img.src = url;
  }
  $('collage-result-body').style.display = 'flex';
  $('collage-close').style.display = 'inline-block';
  $('collage-download').textContent = isZip ? '⬇ 下载分片 zip' : '⬇ 下载结果';
}

function formatSize(bytes) {
  if (bytes < 1024) return bytes + ' B';
  if (bytes < 1048576) return (bytes / 1024).toFixed(1) + ' KB';
  return (bytes / 1048576).toFixed(2) + ' MB';
}

$('collage-close').addEventListener('click', hideResult);

function hideResult() {
  $('collage-result').classList.remove('show');
  $('collage-result-body').style.display = 'none';
  $('collage-loading').style.display = 'none';
  showError('collage-error', '');
}

/* ==================== 通用请求 ==================== */
async function postForm(url, fd) {
  const resp = await fetch(url, { method: 'POST', body: fd });
  const ct = resp.headers.get('Content-Type') || '';
  if (!resp.ok) {
    let msg = '请求失败 (' + resp.status + ')';
    try {
      const j = await resp.json();
      if (j.error) msg = j.error;
    } catch (_) { /* ignore */ }
    throw new Error(msg);
  }
  if (ct.includes('application/json')) {
    const j = await resp.json();
    if (j.error) throw new Error(j.error);
    return null;
  }
  return resp.blob();
}

/* ==================== 水印页 ==================== */
let wmFile = null, wmLogoFile = null, wmResultBlob = null;

bindUploadZone($('wm-upload'), $('wm-file'), (files) => {
  const f = Array.from(files).find((x) => x.type.startsWith('image/'));
  if (!f) return;
  wmFile = f;
  const zone = $('wm-upload');
  zone.classList.add('has-img');
  zone.innerHTML = `<img src="${URL.createObjectURL(f)}" alt="原图">`;
});

$('wm-logo').addEventListener('change', (e) => {
  wmLogoFile = e.target.files[0] || null;
});

$('wm-opacity').addEventListener('input', () => {
  $('wm-op-val').textContent = $('wm-opacity').value + '%';
});

$('wm-run').addEventListener('click', async () => {
  if (!wmFile) { showError('wm-error', '请先上传原图'); return; }
  const fd = new FormData();
  fd.append('image', wmFile);
  fd.append('text', $('wm-text').value.trim());
  fd.append('font_size', $('wm-fontsize').value || '48');
  fd.append('color', $('wm-color').value);
  fd.append('opacity', $('wm-opacity').value);
  fd.append('rotation', $('wm-rotation').value || '0');
  fd.append('position', $('wm-position').value);
  if (wmLogoFile) {
    fd.append('logo', wmLogoFile);
    fd.append('logo_size', $('wm-logosize').value || '200');
    fd.append('logo_opacity', $('wm-logoopacity').value || '100');
  }
  fd.append('format', 'png');

  $('wm-loading').style.display = 'block';
  $('wm-result-img').style.display = 'none';
  $('wm-empty-hint').style.display = 'none';
  $('wm-download').style.display = 'none';
  showError('wm-error', '');

  try {
    const blob = await postForm('/api/watermark', fd);
    if (!blob) return;
    wmResultBlob = blob;
    const img = $('wm-result-img');
    img.src = URL.createObjectURL(blob);
    img.style.display = 'block';
    $('wm-download').style.display = 'inline-block';
  } catch (e) {
    showError('wm-error', e.message);
  } finally {
    $('wm-loading').style.display = 'none';
  }
});

$('wm-download').addEventListener('click', () => {
  if (wmResultBlob) downloadBlob(wmResultBlob, 'watermarked.png');
});

/* ==================== 图片修改页 ==================== */
let edFile = null, edResultBlob = null;

bindUploadZone($('ed-upload'), $('ed-file'), (files) => {
  const f = Array.from(files).find((x) => x.type.startsWith('image/'));
  if (!f) return;
  edFile = f;
  const zone = $('ed-upload');
  zone.classList.add('has-img');
  zone.innerHTML = `<img src="${URL.createObjectURL(f)}" alt="原图">`;
});

['ed-brightness', 'ed-contrast', 'ed-saturation'].forEach((id) => {
  $(id).addEventListener('input', () => {
    const map = { 'ed-brightness': 'ed-br-val', 'ed-contrast': 'ed-co-val', 'ed-saturation': 'ed-sa-val' };
    $(map[id]).textContent = $(id).value;
  });
});

$('ed-run').addEventListener('click', async () => {
  if (!edFile) { showError('ed-error', '请先上传图片'); return; }
  const ops = {
    width: +$('ed-width').value || 0,
    height: +$('ed-height').value || 0,
    scale: +$('ed-scale').value || 0,
    rotate: +$('ed-rotate').value || 0,
    crop_aspect: $('ed-aspect').value,
    grayscale: $('ed-gray').checked,
    invert: $('ed-invert').checked,
    brightness: +$('ed-brightness').value,
    contrast: +$('ed-contrast').value,
    saturation: +$('ed-saturation').value,
    format: $('ed-format').value,
    quality: +$('ed-quality').value || 90,
  };
  const fd = new FormData();
  fd.append('image', edFile);
  fd.append('ops', JSON.stringify(ops));

  $('ed-loading').style.display = 'block';
  $('ed-result-img').style.display = 'none';
  $('ed-empty-hint').style.display = 'none';
  $('ed-download').style.display = 'none';
  showError('ed-error', '');

  try {
    const blob = await postForm('/api/edit', fd);
    if (!blob) return;
    edResultBlob = blob;
    const img = $('ed-result-img');
    img.src = URL.createObjectURL(blob);
    img.style.display = 'block';
    $('ed-download').style.display = 'inline-block';
  } catch (e) {
    showError('ed-error', e.message);
  } finally {
    $('ed-loading').style.display = 'none';
  }
});

$('ed-download').addEventListener('click', () => {
  if (edResultBlob) downloadBlob(edResultBlob, 'edited.' + ($('ed-format').value === 'png' ? 'png' : $('ed-format').value));
});

/* ==================== 支持作者 ==================== */
$('support-btn').addEventListener('click', () => { $('support-modal').style.display = 'flex'; });
$('support-close').addEventListener('click', () => { $('support-modal').style.display = 'none'; });
$('support-modal').addEventListener('click', (e) => {
  if (e.target === $('support-modal')) $('support-modal').style.display = 'none';
});
document.addEventListener('keydown', (e) => {
  if (e.key === 'Escape') $('support-modal').style.display = 'none';
});

/* ==================== 初始化 ==================== */
buildMenu();
renderLibrary();
