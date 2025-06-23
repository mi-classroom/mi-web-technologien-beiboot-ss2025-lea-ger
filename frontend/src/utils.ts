export type FileItem = {
  filepath: string;
  upload_date: string;
  modification_date: string;
  id: number;
};

const IPTC_TAGS_XML_PATH = '/iptc-tags.xml';

export type IPTCTag = {
  id: string;
  name: string;
  type: string;
  writable: boolean;
  description?: string;
  values?: { id: string; value: string }[];
};

export async function getAllTags(): Promise<IPTCTag[]> {
  if (window.__iptcTags) {
    return window.__iptcTags;
  }

  const xml = await fetch(IPTC_TAGS_XML_PATH).then(r => r.text());
  const parser = new DOMParser();
  const xmlDoc = parser.parseFromString(xml, 'application/xml');
  const tables = xmlDoc.getElementsByTagName('table');
  const tags: IPTCTag[] = [];

  for (const table of Array.from(tables)) {
    const tagList = table.getElementsByTagName('tag');
    for (const tag of Array.from(tagList)) {
      const writable = tag.getAttribute('writable') === 'true';
      const values = Array.from(tag.getElementsByTagName('key')).map(key => ({
        id: key.getAttribute('id') || '',
        value: key.textContent || ''
      }));
      tags.push({
        id: tag.getAttribute('id') || '',
        name: tag.getAttribute('name') || '',
        type: tag.getAttribute('type') || '',
        writable,
        description: tag.getElementsByTagName('desc')[0]?.textContent || undefined,
        values: values.length > 0 ? values : undefined
      });
    }
  }

  window.__iptcTags = tags;
  return tags;
}

export async function isWritable(tagName: string): Promise<boolean> {
  const tags = await getAllTags();
  const tag = tags.find(t => t.name === tagName);
  return tag ? tag.writable : false;
}

export async function getTagType(tagName: string): Promise<string | undefined> {
  const tags = await getAllTags();
  const tag = tags.find(t => t.name === tagName);
  return tag?.type;
}

export async function getTagOptions(tagName: string): Promise<{ id: string; value: string }[] | undefined> {
  const tags = await getAllTags();
  const tag = tags.find(t => t.name === tagName);
  return tag?.values;
}