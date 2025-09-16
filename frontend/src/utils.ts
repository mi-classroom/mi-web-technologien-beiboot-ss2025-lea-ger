export type FileItem = {
  filepath: string;
  upload_date: string;
  modification_date: string;
  id: number;
};

export type Folder = {
  id: number;
  name: string;
  children?: Folder[];
}

const IPTC_TAGS_XML_PATH = '/iptc-tags.xml';

export type IPTCTag = {
  id: string;
  name: string;
  type: string;
  g2: string;
  writable: boolean;
  description?: string;
  values?: { id: string; value: string }[];
  relevant?: boolean;
};

export const mostRelevantIPTCTags = [
  {tag: "00", name: "Record Version", description: "Mandatory / The current IPTC Information Interchange Model = 4"},
  {tag: "03", name: "ObjectType", description: "Object Type Reference. 1:News, 2:Data, 3:Advisory x:yy"},
  {
    tag: "04",
    name: "ObjectAttribute",
    description: "Object Attribute Reference. 1:Current, 2:Analysis, 3: Archive .. x:yy"
  },
  {tag: "05", name: "ObjectName", description: "ObjectName. For example \"Skislopes in Norway\" 64 Bytes"},
  {
    tag: "07",
    name: "EditStatus",
    description: "EditStatus. Status of the objectdata accoding to the provider eg \"Correction\" 64 Bytes"
  },
  {tag: "08", name: "EditorialUpdate", description: "Editorial Update - (To a previous object) x:yyy (Num)"},
  {tag: "10", name: "Urgency", description: "Urgency, Editorial urgency 1 = most, 5 = normal, 8 = least x"},
  {
    tag: "12",
    name: "SubjectReference",
    description: "Subject Reference, IPTC:SubRefNum:SubName:SubMatter:SubDetailName 13 to 236 bytes"
  },
  {
    tag: "15",
    name: "Category",
    description: "Category. Identifies the subject of the object in the opinion of the provider AAA (Alpha)"
  },
  {
    tag: "20",
    name: "SupplementalCategory",
    description: "Supplemental Category. Further identifies the subject of the object in the opinion of the provider 32 bytes"
  },
  {
    tag: "22",
    name: "FixtureIdentifier",
    description: "Fixture Identifier. Identifies frequently occuring object data eg \"Euroweather\" 32 bytes (Alpha)"
  },
  {tag: "25", name: "Keyword", description: "Keywords, 64 bytes"},
  {
    tag: "26",
    name: "LocationCode",
    description: "Content Location Code, 3 char code indicating which country the event took place eg NOR AAA (Alpha)"
  },
  {
    tag: "27",
    name: "LocationName",
    description: "Content Location Name. Name of the country the event took place eg \"Norway\" 64 Bytes"
  },
  {
    tag: "30",
    name: "ReleaseDate",
    description: "Release Date. The earliest date the provider intends the object is to be used Date"
  },
  {
    tag: "35",
    name: "ReleaseTime",
    description: "Release Time. The earliest time the provider intends the object is to be used Time"
  },
  {
    tag: "37",
    name: "ExpirationDate",
    description: "Expiration Date. The latest date the provider intends the object to be used Date"
  },
  {
    tag: "38",
    name: "ExpirationTime",
    description: "Expiration Time. The earliest time the provider intends the object to be used Time"
  },
  {tag: "40", name: "SpecialInstructions", description: "Special Instructions. text 256 Bytes"},
  {
    tag: "42",
    name: "ActionAdvised",
    description: "Action Advised. 2 Digits - provider defined. eg 01 = Kill object xx"
  },
  {
    tag: "45",
    name: "ReferenceService",
    description: "Reference Service. Only used to refer to a previous 1:30 (Reference Service) 10 Bytes (Alpha)"
  },
  {tag: "47", name: "ReferenceDate", description: "Reference Date. Only allowed if 45 used"},
  {tag: "50", name: "ReferenceNumber", description: "Reference Number. Only allowed if 45 used 8 Bytes (Num)"},
  {
    tag: "55",
    name: "DateCreated",
    description: "Date Created. Date of the actual event - not when it was digitised Date"
  },
  {
    tag: "60",
    name: "TimeCreated",
    description: "Time Created. Time of the actual event - not when it was digitised Time"
  },
  {
    tag: "62",
    name: "DigitalCreationDate",
    description: "Digital Creation Date. Date the digital representation of the objectdata was created Date"
  },
  {
    tag: "63",
    name: "DigitalCreationTime",
    description: "Digital Creation Time. Time the digital representation of the objectdata was created Time"
  },
  {
    tag: "65",
    name: "OriginatingProgram",
    description: "Originating Program. For example \"Photoshop\" or \"II2LN\" 32 Bytes"
  },
  {tag: "70", name: "ProgramVersion", description: "Program Version. Only use if 65 used eg \"1.2\" 10 Bytes"},
  {tag: "75", name: "ObjectCycle", description: "Object Cycle. Virtually only used in US for Morning Evening or both"},
  {tag: "80", name: "ByLine", description: "Byline. Photographer, for example \"Your Name\" 32 Bytes"},
  {
    tag: "85",
    name: "ByLineTitle",
    description: "ByLine Title. The title or role of photographer to be credited, for example House photographer, correspondent etc 32 Bytes"
  },
  {tag: "90", name: "City", description: "City. According to the practices of the provider eg \"Oslo\" 32 Bytes"},
  {
    tag: "92",
    name: "SubLocation",
    description: "SubLocation. According to the practices of the provider eg \"Holmenkollen\" 32 Bytes"
  },
  {
    tag: "95",
    name: "Province_State",
    description: "Province State. According to the practices of the provider eg \"Vestfold\", \"Adger\" 32 Bytes"
  },
  {
    tag: "100",
    name: "Country_Primary_LocationCode",
    description: "Country Primary Location Code. The country code of the subject matter eg \"USA\" or \"NOR\" AAA (Alpha)"
  },
  {
    tag: "101",
    name: "Country_Primary_LocationName",
    description: "Country Primary Location Name. The country code of the subject matter eg \"Norway\" 64 Bytes"
  },
  {
    tag: "103",
    name: "OriginalTransmissionReference",
    description: "Original Transmission Reference. A code meaning where the object was transmitted from 32 Bytes"
  },
  {tag: "105", name: "Headline", description: "Headline. Synopsis of the subject matter 256 Bytes"},
  {
    tag: "110",
    name: "Credit",
    description: "Credit. Identifies the provider - not necessarily the owner / creator 32 Bytes"
  },
  {tag: "115", name: "Source", description: "Source. Identifies the original owner / creator 32 Bytes"},
  {tag: "116", name: "CopyrightNotice", description: "Copyright Notice. The providers own copyright notice 128 Bytes"},
  {tag: "118", name: "Contact", description: "Contact. For further info on the object 128 Bytes"},
  {
    tag: "120",
    name: "Caption_Abstract",
    description: "Caption and/or Abstract. Full version / rest of the headline 2000 Bytes"
  },
  {tag: "122", name: "Writer_Editor", description: "Writer and/or Editor. Who wrote the caption 32 Bytes"},
  {tag: "130", name: "ImageType", description: "Image Type. Code representing B/w, Y comp of seps etc.Contact x:A"},
  {
    tag: "131",
    name: "ImageOrientation",
    description: "Image Orientation. P = Portrait, L = landscape, S = Square 1 Byte"
  },
  {
    tag: "135",
    name: "LanguageIdentifier",
    description: "Short code identifying the language. For example US for USA and NO for Norway See Also"
  }
];

export async function getAllTags(): Promise<IPTCTag[]> {
  // Cache the tags in a global variable to avoid multiple fetches
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

      const relevantTag = mostRelevantIPTCTags.find(rt => rt.tag === tag.getAttribute('id'));
      tags.push({
        id: tag.getAttribute('id') || '',
        name: tag.getAttribute('name') || '',
        type: tag.getAttribute('type') || '',
        g2: tag.getAttribute('g2') || '',
        writable,
        description: relevantTag?.description || tag.getElementsByTagName('desc')[0]?.textContent || undefined,
        values: values.length > 0 ? values : undefined,
        relevant: !!relevantTag
      });
    }
  }

  window.__iptcTags = tags;
  return tags;
}

export async function getTagByName(tagName: string): Promise<IPTCTag | undefined> {
  const tags = await getAllTags();
  return tags.find(t => t.name === tagName);
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

export async function loadFilesInFolder(folderId: string | number): Promise<void> {
  try {
    const response = await fetch(`${import.meta.env.VITE_APP_API_URL}/api/metadata?folder_id=${folderId}`);
    if (response.ok) {
      return (await response.json()) || [];
    } else {
      console.error('Fehler beim Laden der Dateien');
    }
  } catch (error) {
    console.error('Fehler beim Laden der Dateien', error);
  }
}